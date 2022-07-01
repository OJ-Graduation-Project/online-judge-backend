package util

const (
	DB_HOST = "127.0.0.1"
	DB_PORT = "27017"
	DB_NAME = "onlineJudgeDB"

	CONTESTS_COLLECTION    = "contests"
	PROBLEMS_COLLECTION    = "problems"
	SUBMISSIONS_COLLECTION = "submissions"
	USERS_COLLECTION       = "users"
	SECURITY_CODE          = `
    #ifndef _SECCOMP_BPF_H_
#define _SECCOMP_BPF_H_

#define _GNU_SOURCE 1
#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <errno.h>
#include <signal.h>
#include <string.h>
#include <unistd.h>
#include <errno.h>
#include <linux/audit.h>
#include <linux/bpf.h>
#include <linux/filter.h>
#include <linux/seccomp.h>
#include <linux/unistd.h>
#include <stddef.h>
#include <stdio.h>
#include <sys/prctl.h>
#include <unistd.h>
#include <unistd.h> 
#include <sys/prctl.h>
#ifndef PR_SET_NO_NEW_PRIVS
# define PR_SET_NO_NEW_PRIVS 38
#endif

#include <linux/unistd.h>
#include <linux/audit.h>
#include <linux/filter.h>
#ifdef HAVE_LINUX_SECCOMP_H
# include <linux/seccomp.h>
#endif
#ifndef SECCOMP_MODE_FILTER
# define SECCOMP_MODE_FILTER	2 /* uses user-supplied filter. */
# define SECCOMP_RET_KILL	0x00000000U /* kill the task immediately */
# define SECCOMP_RET_TRAP	0x00030000U /* disallow and force a SIGSYS */
# define SECCOMP_RET_ALLOW	0x7fff0000U /* allow */
struct seccomp_data {
    int nr;
    __u32 arch;
    __u64 instruction_pointer;
    __u64 args[6];
};
#endif
#ifndef SYS_SECCOMP
# define SYS_SECCOMP 1
#endif

#define syscall_nr (offsetof(struct seccomp_data, nr))
#define arch_nr (offsetof(struct seccomp_data, arch))

#if defined(__i386__)
# define REG_SYSCALL	REG_EAX
# define ARCH_NR	AUDIT_ARCH_I386
#elif defined(__x86_64__)
# define REG_SYSCALL	REG_RAX
# define ARCH_NR	AUDIT_ARCH_X86_64
#else
# warning "Platform does not support seccomp filter yet"
# define REG_SYSCALL	0
# define ARCH_NR	0
#endif

#define VALIDATE_ARCHITECTURE \
	BPF_STMT(BPF_LD+BPF_W+BPF_ABS, arch_nr), \
	BPF_JUMP(BPF_JMP+BPF_JEQ+BPF_K, ARCH_NR, 1, 0), \
	BPF_STMT(BPF_RET+BPF_K, SECCOMP_RET_KILL)

#define EXAMINE_SYSCALL \
	BPF_STMT(BPF_LD+BPF_W+BPF_ABS, syscall_nr)

#define ALLOW_SYSCALL(name) \
	BPF_JUMP(BPF_JMP+BPF_JEQ+BPF_K, __NR_##name, 0, 1), \
	BPF_STMT(BPF_RET+BPF_K, SECCOMP_RET_ALLOW)

#define KILL_PROCESS \
	BPF_STMT(BPF_RET+BPF_K, SECCOMP_RET_KILL)

#endif /* _SECCOMP_BPF_H_ */





static void install_syscall_filter(void)
{
	struct sock_filter filter[] = {
		VALIDATE_ARCHITECTURE,
		EXAMINE_SYSCALL,
		ALLOW_SYSCALL(rt_sigreturn),
#ifdef __NR_sigreturn
		ALLOW_SYSCALL(sigreturn),
#endif
		ALLOW_SYSCALL(exit_group),
		ALLOW_SYSCALL(exit),
		ALLOW_SYSCALL(read),
		ALLOW_SYSCALL(write),
        ALLOW_SYSCALL(close),
        ALLOW_SYSCALL(fstat),
        ALLOW_SYSCALL(lseek),
        ALLOW_SYSCALL(mmap),
        ALLOW_SYSCALL(mprotect),
        ALLOW_SYSCALL(munmap),
        ALLOW_SYSCALL(brk),
        ALLOW_SYSCALL(pread64),
        ALLOW_SYSCALL(access),
        ALLOW_SYSCALL(execve),
        ALLOW_SYSCALL(arch_prctl),
        ALLOW_SYSCALL(openat),
		KILL_PROCESS,
	};
	struct sock_fprog prog = {
		.len = (unsigned short)(sizeof(filter)/sizeof(filter[0])),
		.filter = filter,
	};

    prctl(PR_SET_NO_NEW_PRIVS, 1, 0, 0, 0);
    prctl(PR_SET_SECCOMP, SECCOMP_MODE_FILTER, &prog);
}
void implement();
int main(int argc, char **argv)
{
   install_syscall_filter();
    implement();
   
}`

	//Constants for errors:
	DATABASE_FAILED_CONNECTION = "[LOG] Error couldn't connect to database!"
	QUERY                      = "[LOG] Error in query!"
	CURSOR                     = "[LOG] Error in cursor!"
	PING                       = "[LOG] Error in ping!"
	COOKIE                     = "[LOG] Error in getting Cookie!"
	EMPTY_PROBLEM              = "[LOG] No such Problem exists in the database with name: "
	EMPTY_CONTEST              = "[LOG] No such Contest exists in the database with name: "
	EMPTY_CONTESTS             = "[LOG] No Contests exist in the system yet!"
	EMPTY_USER_PROBLEMS        = "[LOG] No Created Problems for this user yet!"
	EMPTY_USER_SUBMISSIONS     = "[LOG] No Submissions for this user yet!"
	EMPTY_TOPIC_PROBLEMS       = "[LOG] No Problems exist for this topic yet!"

	SUBMISSION_ERROR = "[LOG] Error no such Submission exists!"
	USER_NOT_FOUND   = "[LOG] Error no such User exists!"

	MORE_THAN_ONE_CONTEST    = "[LOG] Error more than one Contest with the same name exist where name is: "
	DECODE_CONTEST_FAILED    = "[LOG] Error couldn't decode Contest!"
	DECODE_PROBLEM_FAILED    = "[LOG] Error couldn't decode Problem!"
	DECODE_TOPIC_FAILED      = "[LOG] Error couldn't decode Topic!"
	DECODE_USER_FAILED       = "[LOG] Error couldn't decode User!"
	DECODE_REGISTER_FAILED   = "[LOG] Error couldn't decode Register in Contest!"
	DECODE_SEARCH_FAILED     = "[LOG] Error couldn't decode Search!"
	DECODE_SUBMISSION_FAILED = "[LOG] Error couldn't decode Submission Request"

	SUBMISSION_ID_FAILED = "[LOG] Error couldn't create id for Submission!"
	USER_ID_FAILED       = "[LOG] Error couldn't create id for User!"
	CONTEST_ID_FAILED    = "[LOG] Error couldn't create id for Contest!"
	PROBLEM_ID_FAILED    = "[LOG] Error couldn't create id for Problem!"

	INSERT_CONTEST_FAILED    = "[LOG] Couldn't insert Contest to database!"
	INSERT_PROBLEM_FAILED    = "[LOG] Couldn't insert Problem to database!"
	INSERT_USER_FAILED       = "[LOG] Couldn't insert User to database!"
	INSERT_SUBMISSION_FAILED = "[LOG] Couldn't insert Submission to database!"

	EMAIL_FROM_COOKIE_FAILED = "[LOG] Error in getting authEmail from Cookie!"
	HASHING_PASSWORD_FAILED  = "[LOG] Error couldn't hash password!"
	USER_FROM_EMAIL_FAILED   = "[LOG] Error couldn't fetch User from database."
	INCORRECT_PASSWORD       = "[LOG] Incorrect Password!"
	USER_ERROR               = "[LOG] Error exists with Cursor and User!"
	FETCH_PROBLEM_ID_FAILED  = "[LOG] Error in fetching the problem!"

	//Constants for logging information:
	//DB INFO
	CREATING_DATABASE_CONNECTION = "[LOG] Creating connection with database"
	DATABASE_SUCCESS_CONNECTION  = "[LOG] Connected to database successfully"
	PING_DATABASE                = "[LOG] Pinging database ..."

	//COOKIE INFO
	GETTING_COOKIE            = "[LOG] Getting Cookie"
	DELETING_COOKIE           = "[LOG] Deleting Cookie"
	EMAIL_FROM_COOKIE         = "[LOG] Getting Email from Cookie to fetch user from db ..."
	EMAIL_FROM_COOKIE_SUCCESS = "[LOG] Getting Email from Cookie successfully"

	//CONTESTS INFO
	FETCH_ALL_CONTESTS         = "[LOG] Fetching all contests from database..."
	RETURNING_ALL_CONTESTS     = "[LOG] Returning all contests successfully in the response"
	FETCH_CONTEST              = "[LOG] Fetching Contest from database where Contest name is: "
	FETCH_CONTEST_PROBLEMS     = "[LOG] Fetching problems from database for Contest:  "
	RETURNING_CONTEST_PROBLEMS = "[LOG] Returning all problems for Contest successfully in the response"
	DECODE_CONTEST             = "[LOG] Decoding Contest ..."
	SAVING_CONTEST_IN_DATABASE = "[LOG] Saving Contest in database"
	CREATE_CONTEST_ID          = "[LOG] Creating id for Contest ..."
	CONTEST_ID_SUCCESS         = "[LOG] id is created successfully for Contest: "
	INSERT_CONTEST             = "[LOG] Inserting Contest in database ..."
	INSERT_CONTEST_SUCCESS     = "[LOG] Contest is inserted successfully in database with name: "
	DECODE_CONTEST_SUCCESS     = "[LOG] Contest is decoded successfully"
	FETCHED_CONTEST_SUCCESS    = "[LOG] Contest is fetched successfully from the database"

	//PROBLEM INFO
	DECODE_PROBLEM         = "[LOG] Decoding Problem ..."
	CREATE_PROBLEM_ID      = "[LOG] Creating id for Problem ..."
	PROBLEM_ID_SUCCESS     = "[LOG] id is created successfully for Problem: "
	INSERT_PROBLEM         = "[LOG] Inserting Problem in database ..."
	DECODE_PROBLEM_SUCCESS = "[LOG] Problem is decoded successfully"
	FETCHING_PROBLEM       = "[LOG] Fetching Problem from database with name: "
	RETURNING_PROBLEM      = "[LOG] Returning Problem successfully in the response"
	INSERT_PROBLEM_SUCCESS = "[LOG] Problem is inserted successfully in database with name: "
	FETCHING_PROBLEM_ID    = "[LOG] Fetching Problem from database"

	//USER INFO
	GETTING_USER               = "[LOG] Getting User from database ..."
	SET_USER_AS_WRITER         = "[LOG] Setting User id as Problem's Writer id"
	DECODE_USER                = "[LOG] Decoding User ..."
	DECODE_USER_SUCCESS        = "[LOG] User is decoded successfully"
	FETCHING_USER_FROM_EMAIL   = "[LOG] Fetch User from database with email: "
	COMPARE_HASH               = "[LOG] Comparing User's Password with hashed Password"
	RETURNING_USER             = "[LOG] Returning User successfully in response"
	CREATE_USER_ID             = "[LOG] Creating id for User ..."
	USER_ID_SUCCESS            = "[LOG] id is created successfully for User"
	FETCHING_USER_PROBLEMS     = "[LOG] Fetching Problems from database where writerID is same as userID"
	RETURNING_USER_PROBLEMS    = "[LOG] Returning Problems from database where writerID is same as userID successfully in the response"
	FETCHING_USER_SUBMISSIONS  = "[LOG] Fetching User Submissions from database"
	RETURNING_USER_SUBMISSIONS = "[LOG] Returning User Submissions from database successfully in the response"
	INSERT_USER                = "[LOG] Inserting User in database ..."
	INSERT_USER_SUCCESS        = "[LOG] User is inserted successfully in database with email: "

	//TOPIC INFO
	DECODE_TOPIC                 = "[LOG] Decoding Topic ..."
	DECODE_TOPIC_SUCCESS         = "[LOG] Topic is decoded successfully"
	FETCHING_PROBLEMS_FROM_TOPIC = "[LOG] Fetching Problems from database having topic: "
	RETURNING_DESIRED_PROBLEMS   = "[LOG] Returning desired Problems successfully in response"

	//LOGIN/OUT INFO
	HASHING_PASSWORD         = "[LOG] Hashing password ..."
	HASHING_PASSWORD_SUCCESS = "[LOG] Hashing password is done successfully"
	LOGGED_OUT               = "[LOG] User logged out successfully"

	//REGISTER INFO
	DECODE_REGISTER          = "[LOG] Decoding Register in Contest ..."
	DECODE_REGISTER_SUCCESS  = "[LOG] Register in Contest is decoded successfully"
	UPDATE_CONTEST_WITH_USER = "[LOG] Update Contest with new Registered User"
	UPDATE_USER_WITH_CONTEST = "[LOG] Update User with new Contest"

	//SEARCH INFO
	DECODE_SEARCH            = "[LOG] Decoding Search ..."
	DECODE_SEARCH_SUCCESS    = "[LOG] Search is decoded successfully"
	FETCHING_SEARCH_PROBLEMS = "[LOG] Fetching Problems from database which contain substring: "

	//SUBMISSION REQUEST INFO
	DECODE_SUBMISSION           = "[LOG] Decoding Submission Request ..."
	DECODE_SUBMISSION_SUCCESS   = "[LOG] Submission Request is decoded successfully"
	COMPILE                     = "[LOG] Compiling and Running the Submission ..."
	NOT_CONTEST_AND_WRONG       = "[LOG] Submission isn't in contest and is wrong submission"
	CONTEST_AND_WRONG           = "[LOG] Submission is in contest and is wrong submission"
	CONTEST_AND_CORRECT         = "[LOG] Submission is in contest and is accepted submission"
	CREATE_SUBMISSION_ID        = "[LOG] Creating id for Submission ..."
	SUBMISSION_ID_SUCCESS       = "[LOG] id is created successfully for Submission"
	INSERT_SUBMISSION           = "[LOG] Inserting Submission in database ..."
	INSERT_SUBMISSION_SUCCESS   = "[LOG] Submission is inserted successfully in database"
	FETCH_SUBMISSION_OF_PROBLEM = "[LOG] Fetching Submssion of Problem: "
	RETURN_SUBMISSION           = "[LOG] Return Submission successfully in response"
)
