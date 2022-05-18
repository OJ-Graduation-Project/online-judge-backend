package util

const (
    DB_HOST = "127.0.0.1"
    DB_PORT = "27017"
    DB_NAME = "onlineJudgeDB"

    CONTESTS_COLLECTION    = "contests"
    PROBLEMS_COLLECTION    = "problems"
    SUBMISSIONS_COLLECTION = "submissions"
    USERS_COLLECTION       = "users"


    //Constants for errors:
    DATABASE_FAILED_CONNECTION      = "[LOG] Error couldn't connect to database!"
    QUERY                           = "[LOG] Error in query!"
    CURSOR                          = "[LOG] Error in cursor!"
    PING                            = "[LOG] Error in ping!"
    COOKIE                          = "[LOG] Error in getting Cookie!"
    EMPTY_PROBLEM                   = "[LOG] No such Problem exists in the database with name: "
    EMPTY_CONTEST                   = "[LOG] No such Contest exists in the database with name: "
    SUBMISSION_ERROR                = "[LOG] Error no such Submission exists!"
    USER_NOT_FOUND                  = "[LOG] Error no such User exists!"

    MORE_THAN_ONE_CONTEST           = "[LOG] Error more than one Contest with the same name exist where name is: "
    DECODE_CONTEST_FAILED           = "[LOG] Error couldn't decode Contest!"
    DECODE_PROBLEM_FAILED           = "[LOG] Error couldn't decode Problem!"
    DECODE_TOPIC_FAILED             = "[LOG] Error couldn't decode Topic!"
    DECODE_USER_FAILED              = "[LOG] Error couldn't decode User!"
    DECODE_REGISTER_FAILED          = "[LOG] Error couldn't decode Register in Contest!"
    DECODE_SEARCH_FAILED            = "[LOG] Error couldn't decode Search!"
    DECODE_SUBMISSION_FAILED        = "[LOG] Error couldn't decode Submission Request"
    
    SUBMISSION_ID_FAILED            = "[LOG] Error couldn't create id for Submission!"
    USER_ID_FAILED                  = "[LOG] Error couldn't create id for User!"
    CONTEST_ID_FAILED               = "[LOG] Error couldn't create id for Contest!"
    PROBLEM_ID_FAILED               = "[LOG] Error couldn't create id for Problem!"

    INSERT_CONTEST_FAILED           = "[LOG] Couldn't insert Contest to database!"
    INSERT_PROBLEM_FAILED           = "[LOG] Couldn't insert Problem to database!"
    INSERT_USER_FAILED              = "[LOG] Couldn't insert User to database!"
    INSERT_SUBMISSION_FAILED        = "[LOG] Couldn't insert Submission to database!"

    EMAIL_FROM_COOKIE_FAILED        = "[LOG] Error in getting authEmail from Cookie!"
    HASHING_PASSWORD_FAILED         = "[LOG] Error couldn't hash password!"
    USER_FROM_EMAIL_FAILED          = "[LOG] Error couldn't fetch User from database."
    INCORRECT_PASSWORD              = "[LOG] Incorrect Password!"
    USER_ERROR                      = "[LOG] Error exists with Cursor and User!"
    FETCH_PROBLEM_ID_FAILED         = "[LOG] Error in fetching the problem!"



    //Constants for logging information:
    //DB INFO
    CREATING_DATABASE_CONNECTION    = "[LOG] Creating connection with database"
    DATABASE_SUCCESS_CONNECTION     = "[LOG] Connected to database successfully"
    PING_DATABASE                   = "[LOG] Pinging database ..."
    
    //COOKIE INFO
    GETTING_COOKIE                  = "[LOG] Getting Cookie"
    DELETING_COOKIE                 = "[LOG] Deleting Cookie"
    EMAIL_FROM_COOKIE               = "[LOG] Getting Email from Cookie to fetch user from db ..."
    EMAIL_FROM_COOKIE_SUCCESS       = "[LOG] Getting Email from Cookie successfully"
    
    //CONTESTS INFO
    FETCH_ALL_CONTESTS              = "[LOG] Fetching all contests from database..."
    RETURNING_ALL_CONTESTS          = "[LOG] Returning all contests successfully in the response"
    FETCH_CONTEST                   = "[LOG] Fetching Contest from database where Contest name is: "
    FETCH_CONTEST_PROBLEMS          = "[LOG] Fetching problems from database for Contest:  "
    RETURNING_CONTEST_PROBLEMS      = "[LOG] Returning all problems for Contest successfully in the response"
    DECODE_CONTEST                  = "[LOG] Decoding Contest ..."
    SAVING_CONTEST_IN_DATABASE      = "[LOG] Saving Contest in database"
    CREATE_CONTEST_ID               = "[LOG] Creating id for Contest ..."
    CONTEST_ID_SUCCESS              = "[LOG] id is created successfully for Contest: "
    INSERT_CONTEST                  = "[LOG] Inserting Contest in database ..."
    INSERT_CONTEST_SUCCESS          = "[LOG] Contest is inserted successfully in database with name: "
    DECODE_CONTEST_SUCCESS          = "[LOG] Contest is decoded successfully"
    FETCHED_CONTEST_SUCCESS         = "[LOG] Contest is fetched successfully from the database"

    //PROBLEM INFO
    DECODE_PROBLEM                  = "[LOG] Decoding Problem ..."
    CREATE_PROBLEM_ID               = "[LOG] Creating id for Problem ..."
    PROBLEM_ID_SUCCESS              = "[LOG] id is created successfully for Problem: "
    INSERT_PROBLEM                  = "[LOG] Inserting Problem in database ..."
    DECODE_PROBLEM_SUCCESS          = "[LOG] Problem is decoded successfully"
    FETCHING_PROBLEM                = "[LOG] Fetching Problem from database with name: "
    RETURNING_PROBLEM               = "[LOG] Returning Problem successfully in the response"
    INSERT_PROBLEM_SUCCESS          = "[LOG] Problem is inserted successfully in database with name: "
    FETCHING_PROBLEM_ID             = "[LOG] Fetching Problem from database"

    //USER INFO
    GETTING_USER                    = "[LOG] Getting User from database ..."
    SET_USER_AS_WRITER              = "[LOG] Setting User id as Problem's Writer id"
    DECODE_USER                     = "[LOG] Decoding User ..."
    DECODE_USER_SUCCESS             = "[LOG] User is decoded successfully"
    FETCHING_USER_FROM_EMAIL        = "[LOG] Fetch User from database with email: "
    COMPARE_HASH                    = "[LOG] Comparing User's Password with hashed Password"
    RETURNING_USER                  = "[LOG] Returning User successfully in response"
    CREATE_USER_ID                  = "[LOG] Creating id for User ..."
    USER_ID_SUCCESS                 = "[LOG] id is created successfully for User"
    FETCHING_USER_PROBLEMS          = "[LOG] Fetching Problems from database where writerID is same as userID"
    RETURNING_USER_PROBLEMS         = "[LOG] Returning Problems from database where writerID is same as userID successfully in the response"
    FETCHING_USER_SUBMISSIONS       = "[LOG] Fetching User Submissions from database"
    RETURNING_USER_SUBMISSIONS      = "[LOG] Returning User Submissions from database successfully in the response"
    INSERT_USER                     = "[LOG] Inserting User in database ..."
    INSERT_USER_SUCCESS             = "[LOG] User is inserted successfully in database with email: "

    //TOPIC INFO
    DECODE_TOPIC                    = "[LOG] Decoding Topic ..."
    DECODE_TOPIC_SUCCESS            = "[LOG] Topic is decoded successfully"
    FETCHING_PROBLEMS_FROM_TOPIC    = "[LOG] Fetching Problems from database having topic: "
    RETURNING_DESIRED_PROBLEMS      = "[LOG] Returning desired Problems successfully in response"

    //LOGIN/OUT INFO
    HASHING_PASSWORD                = "[LOG] Hashing password ..."
    HASHING_PASSWORD_SUCCESS        = "[LOG] Hashing password is done successfully"
    LOGGED_OUT                      = "[LOG] User logged out successfully"

    //REGISTER INFO
    DECODE_REGISTER                 = "[LOG] Decoding Register in Contest ..."
    DECODE_REGISTER_SUCCESS         = "[LOG] Register in Contest is decoded successfully"
    UPDATE_CONTEST_WITH_USER        = "[LOG] Update Contest with new Registered User"
    UPDATE_USER_WITH_CONTEST        = "[LOG] Update User with new Contest"

    //SEARCH INFO
    DECODE_SEARCH                   = "[LOG] Decoding Search ..."
    DECODE_SEARCH_SUCCESS           = "[LOG] Search is decoded successfully"
    FETCHING_SEARCH_PROBLEMS        = "[LOG] Fetching Problems from database which contain substring: "

    //SUBMISSION REQUEST INFO
    DECODE_SUBMISSION               = "[LOG] Decoding Submission Request ..."
    DECODE_SUBMISSION_SUCCESS       = "[LOG] Submission Request is decoded successfully"
    COMPILE                         = "[LOG] Compiling and Running the Submission ..."
    NOT_CONTEST_AND_WRONG           = "[LOG] Submission isn't in contest and is wrong submission"
    CONTEST_AND_WRONG               = "[LOG] Submission is in contest and is wrong submission"
    CONTEST_AND_CORRECT             = "[LOG] Submission is in contest and is accepted submission"
    CREATE_SUBMISSION_ID            = "[LOG] Creating id for Submission ..."
    SUBMISSION_ID_SUCCESS           = "[LOG] id is created successfully for Submission"
    INSERT_SUBMISSION               = "[LOG] Inserting Submission in database ..."
    INSERT_SUBMISSION_SUCCESS       = "[LOG] Submission is inserted successfully in database"
    FETCH_SUBMISSION_OF_PROBLEM     = "[LOG] Fetching Submssion of Problem: "
    RETURN_SUBMISSION               = "[LOG] Return Submission successfully in response"

)

