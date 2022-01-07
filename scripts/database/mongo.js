const HOST = '127.0.0.1'
const PORT = '27017'
const DB_NAME = 'onlineJudgeDB'

const BASE = './scripts/database/data'

const CONTESTS = 'contests.json'
const PROBLEMS = 'problems.json'
const SUBMISSIONS = 'submissions.json'
const USERS = 'users.json'

var db = connect(`${HOST}:${PORT}/${DB_NAME}`),
    onlineJudge = null;


db.contests.insert(JSON.parse(cat(`${BASE}/${CONTESTS}`)))
db.problems.insert(JSON.parse(cat(`${BASE}/${PROBLEMS}`)))
db.submissions.insert(JSON.parse(cat(`${BASE}/${SUBMISSIONS}`)))
db.users.insert(JSON.parse(cat(`${BASE}/${USERS}`)))

