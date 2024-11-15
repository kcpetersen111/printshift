"Requirments":

Admin account - Creates other professor accounts. This goes to Dr. Dave

Professor accounts can assign students to printers

Scheduler modes:
  - Assign mode: Printer is open for assignment (professor assigns)
  - Open mode: Printer is open for anyone to sign up, regardless of class status (desperate times call for desperate measures)

Name and desc for each printer

Students are assigned to classes - 1 printer per clsas

# API Spec

## Endpoints

CreateUser:
```
{
  "Name": String,
  "AccessLevel": Int(enum),
  "Email": String,
  "Classes": String[]
}
```
GetUser

UpdateUser

DeleteUser
## Models
User:
```
{
    name: string,
    accessLevel: AccessLevel,
    email: string,
    classes: string[],
    printers: string[],
    printersCanAssign: number
}
```

## Enums
AccessLevel:
```
{
    Admin = 1,
    Professor = 2,
    Student = 3
}
```