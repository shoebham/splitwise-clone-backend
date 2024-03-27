entities:
	- transactions (category type, can be a column)
	- users
	- groups

flow:
	- user adds/delete/edit transactions
	- user creates groups
	- user sends money
	- user settles transactions
	- user sees analytics

users:
	uid int (pk)
	name string
	groups int (fk)
	balance int (+/-)

groups:
	gid
	name
	members (csv field)

transactions:
	tid int
	name string 
	amount int
	type (equal, unequal, etc)
	members string (csv field) (shubham, vinod )
	timestamp

one user <-> multi groups
one user <-> multi transactions
one user <-> multi users
one user -> one user (without group)

userTransactions:
	uid int (the person who is owed) (from in request)
	gid int
	tid int
	owesId int (person who owes)
	amount int 

userGroups:
	uid int 
	gid int
	total int (+/-)

usersToUsers:
	reciever int 
	sender int
	total int (+/-)

endpoints:
	- /addTransaction
		req body: 
		{ 
			from: "",
			title:"",
			type : "",
			with:["name":"amount"],
			group:"",
			addedBy:"", 
		}
		res body:
		{
			owed:" from in req",
			owes:["vinod":"100"],["shubham":"100"]
		}
		after request is recvd
		updateOwedUser()
		updateOwesUser()

	- /updateUser
		req 
		{

		}

extra effort: payment integration

