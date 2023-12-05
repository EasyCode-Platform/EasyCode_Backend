package storage

import (
	"database/sql"
	"github.com/EasyCode-Platform/EasyCode_Backend/src/model"
	"github.com/google/uuid"

	"log"
)

type PostgresqlStorage struct {
	db *sql.DB
}

// NewPostgresqlStorage 创建一个新的 PostgreSQL 存储实例
func NewPostgresqlStorage(db *sql.DB) *PostgresqlStorage {
	return &PostgresqlStorage{db}
}

// InsertApp 插入一个应用数据到数据库
func (ps *PostgresqlStorage) InsertApp(appData model.AppData) error {
	// 准备 SQL 语句，注意 PostgreSQL 使用 $1, $2 等作为占位符
	query := "INSERT INTO app_data (aid, name) VALUES ($1, $2)"
	// 执行 SQL 语句
	_, err := ps.db.Exec(query, appData.Aid, appData.Name)
	if err != nil {
		log.Println("Error inserting app:", err)
		return err
	}
	return nil
}

// InsertTable 插入一个数据表数据到数据库
func (ps *PostgresqlStorage) InsertTable(tableData model.Table) error {
	// 准备 SQL 语句，注意 PostgreSQL 使用 $1, $2 等作为占位符
	query := "INSERT INTO tables (tid, name, app_aid) VALUES ($1, $2, $3)"
	// 执行 SQL 语句
	_, err := ps.db.Exec(query, tableData.Tid, tableData.Name, tableData.AppAid)
	if err != nil {
		log.Println("Error inserting table:", err)
		return err
	}
	return nil
}

func (ps *PostgresqlStorage) GetAppsData() ([]model.AppData, error) {
	// 查询以获取应用数据和关联的表信息
	query := `
    SELECT ad.aid, ad.name, t.tid, t.name, t.app_aid
    FROM app_data ad
    LEFT JOIN tables t ON ad.aid = t.app_aid
    ORDER BY ad.aid, t.tid
    `
	rows, err := ps.db.Query(query)
	if err != nil {
		log.Println("Error querying apps with tables:", err)
		return nil, err
	}
	defer rows.Close()

	var apps []model.AppData
	var currentApp *model.AppData

	// 遍历查询结果
	for rows.Next() {
		var aid uuid.UUID
		var appName string
		var tidStr sql.NullString // 使用 sql.NullString 来接收可能的 NULL 值
		var tableName sql.NullString
		var appAid uuid.UUID // 新增：用于存储从表中获取的 app_aid

		err := rows.Scan(&aid, &appName, &tidStr, &tableName, &appAid)
		if err != nil {
			log.Println("Error scanning app with tables:", err)
			return nil, err
		}

		if currentApp == nil || currentApp.Aid != aid {
			if currentApp != nil {
				apps = append(apps, *currentApp)
			}
			currentApp = &model.AppData{
				Aid:    aid,
				Name:   appName,
				Tables: []model.Table{}, // 初始化 Tables 切片
			}
		}

		var table model.Table // 初始化 table 变量
		table.AppAid = appAid // 设置 table 的 Aid 字段

		if tidStr.Valid {
			tidUUID, err := uuid.Parse(tidStr.String)
			if err != nil {
				log.Println("Error parsing TID UUID:", err)
				continue
			}
			table.Tid = tidUUID
		}

		if tableName.Valid {
			table.Name = tableName.String
			currentApp.Tables = append(currentApp.Tables, table)
		}
	}

	// 添加最后一个应用
	if currentApp != nil {
		apps = append(apps, *currentApp)
	}

	return apps, nil
}

func (ps *PostgresqlStorage) CreateNewTable(aid uuid.UUID) (*model.Table, error) {
	// 生成一个新的 Table ID (tid)，使用 UUID/NanoID
	tid := uuid.New()

	// 定义创建新表的 SQL 语句
	query := `
        INSERT INTO tables (tid, name, app_aid)
        VALUES ($1, $2, $3)
        RETURNING tid, name
    `

	// 执行 SQL 语句
	var table model.Table
	err := ps.db.QueryRow(query, tid, "未命名表单", aid).Scan(&table.Tid, &table.Name)
	if err != nil {
		log.Println("Error creating new table:", err)
		return nil, err
	}

	// 返回新创建的表
	return &table, nil
}

func (ps *PostgresqlStorage) RenameTable(tid uuid.UUID, newName string) (*model.Table, error) {
	// 定义更新表名的 SQL 语句
	updateQuery := `
    	UPDATE tables
        SET name = $1
		WHERE tid = $2;`

	// 执行更新语句
	_, err := ps.db.Exec(updateQuery, newName, tid)
	if err != nil {
		log.Println("Error updating table name:", err)
		return nil, err
	}

	// 定义获取更新后的表信息的 SQL 语句
	selectQuery := `
		SELECT tid, name
		FROM tables
		WHERE tid = $1;`

	// 执行查询语句
	var table model.Table
	err = ps.db.QueryRow(selectQuery, tid).Scan(&table.Tid, &table.Name)
	if err != nil {
		log.Println("Error fetching updated table:", err)
		return nil, err
	}

	// 返回重命名的表
	return &table, nil
}

func (ps *PostgresqlStorage) DeleteTable(tid uuid.UUID) error {
	// 定义删除表的 SQL 语句
	deleteQuery := `
		DELETE FROM tables
		WHERE tid = $1;
	`

	// 执行删除语句
	_, err := ps.db.Exec(deleteQuery, tid)
	if err != nil {
		log.Println("Error deleting table:", err)
		return err
	}

	// 返回 nil 表示成功删除
	return nil
}

func (ps *PostgresqlStorage) GetTableData(tid uuid.UUID) (model.TableData, error) {
	// 查询字段数据
	fieldQuery := `
        SELECT name, type
        FROM table_fields
        WHERE tid = $1
        ORDER BY name
    `
	fieldRows, err := ps.db.Query(fieldQuery, tid)
	if err != nil {
		log.Println("Error querying table fields:", err)
		return model.TableData{}, err
	}
	defer fieldRows.Close()

	var fields []model.Field

	for fieldRows.Next() {
		var name string
		var fieldType string

		err := fieldRows.Scan(&name, &fieldType)
		if err != nil {
			log.Println("Error scanning table fields:", err)
			return model.TableData{}, err
		}

		field := model.Field{
			Name: name,
			Type: fieldType,
		}

		fields = append(fields, field)
	}

	// 查询记录数据
	recordQuery := `
        SELECT entity_id, field_name, field_value
        FROM table_records
        WHERE tid = $1
        ORDER BY entity_id, field_name
    `
	recordRows, err := ps.db.Query(recordQuery, tid)
	if err != nil {
		log.Println("Error querying table records:", err)
		return model.TableData{}, err
	}
	defer recordRows.Close()

	tempRecords := make(map[int]map[string]interface{})
	var records []map[string]interface{}

	for recordRows.Next() {
		var entityID int
		var fieldName string
		var fieldValue interface{} // 使用 interface{} 来处理不同类型的值

		err := recordRows.Scan(&entityID, &fieldName, &fieldValue)
		if err != nil {
			log.Println("Error scanning table records:", err)
			return model.TableData{}, err
		}

		// 检查是否已存在该 entityID 的记录
		if _, exists := tempRecords[entityID]; !exists {
			tempRecords[entityID] = make(map[string]interface{})
		}

		// 添加字段到对应的记录中
		tempRecords[entityID][fieldName] = fieldValue
	}

	// 将临时 map 转换为所需的 records 列表
	for _, record := range tempRecords {
		records = append(records, record)
	}

	tableData := model.TableData{
		Fields:  fields,
		Records: records,
	}

	return tableData, nil
}
