package config

const SERVER_API_FULL_URL = "http://localhost:5000/"
const SERVER_API_DOMAIN = "localhost:5000"

const CDN_IMAGE_SERVER_DOMAIN = "localhost:5000"

//const CDN_IMAGE_SERVER_FULL_URL = "http://localhost:5000/upload/"
const CDN_IMAGE_SERVER_FULL_URL = "http://192.168.0.10:5000/upload/"

//const CDN_CHAT_MSG_UPLOAD_URL = "http://localhost:5000/"
const CDN_CHAT_MSG_UPLOAD_URL = "http://192.168.0.10:5000/"

const NUMBER_OF_USER_AVATAR_DIRS = 5

const IS_DEBUG = true
const DEBUG_LOG_SQL = IS_DEBUG
const IS_PRODUCTION = false
const DEBUG_DELAY_RUN_STARTUPS = false //IS_DEBUG

const UPLOAD_DIR = "./upload/"
const UPLOAD_DIR_POSTS = UPLOAD_DIR + ""

const BUCKET_SIZE = 30

const TAGS_RELOAD_TOP_INTERVAL_MINS = 60
