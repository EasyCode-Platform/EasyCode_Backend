set -e

mongo <<EOF
db = db.getSiblingDB('ec-backend')
db.createCollection('components')


EOF
