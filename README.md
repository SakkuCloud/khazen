# KHAZEN - A web service to work with SQL

[![N|Solid](https://avatars3.githubusercontent.com/u/44247427?s=200&v=4)](https://github.com/SakkuCloud)

Khazen is a web service that manage SQL databases such as MySQL.

# Motivations
In SAKKU team we have several modules that needs to make root privileged actions in database. So we need a manager to listen in endpoints and make this actions in database host to prevent issues like:
- Using ROOT username and password in every module.
- ROOT user can login out of databse localhost.
- No control in actions that every module can do with ROOT user.