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
	DATABASE_FAILED_CONNECTION		= "Error couldn't connect to database!"
	QUERY							= "Error in query!"
	CURSOR 							= "Error in cursor!"
	PING 							= "Error in ping!"
	COOKIE 							= "Error in getting Cookie!"
	EMPTY_PROBLEM					= "No such Problem exists in the database with name: "
	EMPTY_CONTEST					= "No such Contest exists in the database with name: "
	SUBMISSION_ERROR				= "Error no such Submission exists!"
	USER_NOT_FOUND					= "Error no such User exists!"

	MORE_THAN_ONE_CONTEST			= "Error more than one Contest with the same name exist where name is: "
	DECODE_CONTEST_FAILED 			= "Error couldn't decode Contest!"
	DECODE_PROBLEM_FAILED 			= "Error couldn't decode Problem!"
	DECODE_TOPIC_FAILED 			= "Error couldn't decode Topic!"
	DECODE_USER_FAILED	 			= "Error couldn't decode User!"
	DECODE_REGISTER_FAILED		= "Error couldn't decode Register in Contest!"
	DECODE_SEARCH_FAILED			= "Error couldn't decode Search!"
	DECODE_SUBMISSION_FAILED		= "Error couldn't decode Submission Request"
	
	SUBMISSION_ID_FAILED			= "Error couldn't create id for Submission!"
	USER_ID_FAILED					= "Error couldn't create id for User!"
	CONTEST_ID_FAILED				= "Error couldn't create id for Contest!"
	PROBLEM_ID_FAILED				= "Error couldn't create id for Problem!"

	INSERT_CONTEST_FAILED			= "Couldn't insert Contest to database!"
	INSERT_PROBLEM_FAILED			= "Couldn't insert Problem to database!"
	INSERT_USER_FAILED				= "Couldn't insert User to database!"
	INSERT_SUBMISSION_FAILED		= "Couldn't insert Submission to database!"

	EMAIL_FROM_COOKIE_FAILED 		= "Error in getting authEmail from Cookie!"
	HASHING_PASSWORD_FAILED			= "Error couldn't hash password!"
	USER_FROM_EMAIL_FAILED 			= "Error couldn't fetch User from database."
	INCORRECT_PASSWORD				= "Incorrect Password!"
	USER_ERROR						= "Error exists with Cursor and User!"
	FETCH_PROBLEM_ID_FAILED 		= "Error in fetching the problem!"



	//Constants for logging information:
	//DB INFO
	CREATING_DATABASE_CONNECTION 	= "Creating connection with database"
	DATABASE_SUCCESS_CONNECTION		= "Connected to database successfully"
	PING_DATABASE					= "Pinging database ..."
	
	//COOKIE INFO
	GETTING_COOKIE 					= "Getting Cookie"
	DELETING_COOKIE					= "Deleting Cookie"
	EMAIL_FROM_COOKIE				= "Getting Email from Cookie to fetch user from db ..."
	EMAIL_FROM_COOKIE_SUCCESS		= "Getting Email from Cookie successfully"
	
	//CONTESTS INFO
	FETCH_ALL_CONTESTS				= "Fetching all contests from database..."
	RETURNING_ALL_CONTESTS			= "Returning all contests successfully in the response"
	FETCH_CONTEST					= "Fetching Contest from database where Contest name is: "
	FETCH_CONTEST_PROBLEMS			= "Fetching problems from database for Contest:  "
	RETURNING_CONTEST_PROBLEMS		= "Returning all problems for Contest successfully in the response"
	DECODE_CONTEST					= "Decoding Contest ..."
	SAVING_CONTEST_IN_DATABASE		= "Saving Contest in database"
	CREATE_CONTEST_ID				= "Creating id for Contest ..."
	CONTEST_ID_SUCCESS				= "id is created successfully for Contest: "
	INSERT_CONTEST 					= "Inserting Contest in database ..."
	INSERT_CONTEST_SUCCESS			= "Contest is inserted successfully in database with name: "
	DECODE_CONTEST_SUCCESS 			= "Contest is decoded successfully"
	FETCHED_CONTEST_SUCCESS			= "Contest is fetched successfully from the database"

	//PROBLEM INFO
	DECODE_PROBLEM					= "Decoding Problem ..."
	CREATE_PROBLEM_ID				= "Creating id for Problem ..."
	PROBLEM_ID_SUCCESS				= "id is created successfully for Problem: "
	INSERT_PROBLEM 					= "Inserting Problem in database ..."
	DECODE_PROBLEM_SUCCESS 			= "Problem is decoded successfully"
	FETCHING_PROBLEM				= "Fetching Problem from database with name: "
	RETURNING_PROBLEM				= "Returning Problem successfully in the response"
	INSERT_PROBLEM_SUCCESS			= "Problem is inserted successfully in database with name: "
	FETCHING_PROBLEM_ID				= "Fetching Problem from database"

	//USER INFO
	GETTING_USER					= "Getting User from database ..."
	SET_USER_AS_WRITER				= "Setting User id as Problem's Writer id"
	DECODE_USER						= "Decoding User ..."
	DECODE_USER_SUCCESS 			= "User is decoded successfully"
	FETCHING_USER_FROM_EMAIL 		= "Fetch User from database with email: "
	COMPARE_HASH					= "Comparing User's Password with hashed Password"
	RETURNING_USER					= "Returning User successfully in response"
	CREATE_USER_ID					= "Creating id for User ..."
	USER_ID_SUCCESS					= "id is created successfully for User"
	FETCHING_USER_PROBLEMS			= "Fetching Problems from database where writerID is same as userID"
	RETURNING_USER_PROBLEMS			= "Returning Problems from database where writerID is same as userID successfully in the response"
	FETCHING_USER_SUBMISSIONS		= "Fetching User Submissions from database"
	RETURNING_USER_SUBMISSIONS		= "Returning User Submissions from database successfully in the response"
	INSERT_USER 					= "Inserting User in database ..."
	INSERT_USER_SUCCESS				= "User is inserted successfully in database with email: "

	//TOPIC INFO
	DECODE_TOPIC					= "Decoding Topic ..."
	DECODE_TOPIC_SUCCESS 			= "Topic is decoded successfully"
	FETCHING_PROBLEMS_FROM_TOPIC 	= "Fetching Problems from database having topic: "
	RETURNING_DESIRED_PROBLEMS		= "Returning desired Problems successfully in response"

	//LOGIN/OUT INFO
	HASHING_PASSWORD				= "Hashing password ..."
	HASHING_PASSWORD_SUCCESS		= "Hashing password is done successfully"
	LOGGED_OUT						= "User logged out successfully"

	//REGISTER INFO
	DECODE_REGISTER					= "Decoding Register in Contest ..."
	DECODE_REGISTER_SUCCESS			= "Register in Contest is decoded successfully"
	UPDATE_CONTEST_WITH_USER		= "Update Contest with new Registered User"
	UPDATE_USER_WITH_CONTEST		= "Update User with new Contest"

	//SEARCH INFO
	DECODE_SEARCH					= "Decoding Search ..."
	DECODE_SEARCH_SUCCESS			= "Search is decoded successfully"
	FETCHING_SEARCH_PROBLEMS		= "Fetching Problems from database which contain substring: "

	//SUBMISSION REQUEST INFO
	DECODE_SUBMISSION				= "Decoding Submission Request ..."
	DECODE_SUBMISSION_SUCCESS		= "Submission Request is decoded successfully"
	COMPILE							= "Compiling and Running the Submission ..."
	NOT_CONTEST_AND_WRONG			= "Submission isn't in contest and is wrong submission"
	CONTEST_AND_WRONG				= "Submission is in contest and is wrong submission"
	CONTEST_AND_CORRECT				= "Submission is in contest and is accepted submission"
	CREATE_SUBMISSION_ID			= "Creating id for Submission ..."
	SUBMISSION_ID_SUCCESS			= "id is created successfully for Submission"
	INSERT_SUBMISSION 				= "Inserting Submission in database ..."
	INSERT_SUBMISSION_SUCCESS		= "Submission is inserted successfully in database"
	FETCH_SUBMISSION_OF_PROBLEM		= "Fetching Submssion of Problem: "
	RETURN_SUBMISSION 				= "Return Submission successfully in response"



)
