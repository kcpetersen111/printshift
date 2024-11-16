export enum AccessLevel {
    Unknown = 0,
    Admin = 1,
    Professor = 2,
    Student = 3
}

export const convertToStringAccessLevel = (accessLevel: AccessLevel) => {
    switch (accessLevel) {
        case AccessLevel.Admin:
            return "Admin";
        case AccessLevel.Professor:
            return "Professor";
        case AccessLevel.Student:
            return "Student";
        default:
            throw new Error("Access Level unrecognized: " + accessLevel);            
    }
}