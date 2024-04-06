entities: - transactions (category type, can be a column) - users - groups

flow: - user adds/delete/edit transactions - user creates groups - user sends money - user settles transactions - user sees analytics

users: uid int (pk) name string groups int (fk) balance int (+/-)

groups: gid name members (csv field)

transactions: tid int name string amount int type (equal, unequal, etc) members string (csv field) (shubham, vinod ) timestamp

one user <-> multi groups one user <-> multi transactions one user <-> multi users one user -> one user (without group)

userTransactions: uid int (the person who is owed) (from in request) gid int tid int owesId int (person who owes) amount int

userGroups: uid int gid int total int (+/-)

usersToUsers: reciever int sender int total int (+/-)


extra effort: payment integration

Types of Expense:
1. split equally
2. split unequally (Amount Only)

endpoints:

Expense Module:
/expense -> id POST
/expense/:id PUT
/expense/:id DELETE
/expense/:id/settleUp POST

Group Module:
/group -> id POST
/group/:id PUT
/group/addMember -> mid POST
/group/deleteMember/:mid DELETE
/group/:id DELETE
/group/:id/updateTransactions POST//

User Module:
/getGroupMembers


User {
    UserID,
    Name,
    Balance,
    Map<User, double> Owes,
    Map<User, double> Owed,
}


Expense {
    ExpenseID,
    Description,
    Amount,
    Timestamp,
    User Added,
    User Paid,
    Map<User, double> Members,
    isEqually,
    isSettled,
    GroupID
}

Group {
    GroupID,
    GroupName,
    List<User> Members
}
