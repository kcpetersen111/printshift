import { AccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { AdminPanel } from "./AdminPanel";

export const ManagementPanel = (user: User) => {
    let access: string;
    switch (user.accessLevel) {
        case AccessLevel.Admin:
            access = "Admin";
            break;
        case AccessLevel.Professor:
            access = "Professor";
            break;
        case AccessLevel.Student:
            access = "Student";
            break;
        default:
            throw new Error("Access Level unrecognized: " + user.accessLevel);            
    }

    return (
        <div>
            <h1>{access} Panel</h1>
            <AdminPanel user={user} />
        </div>
    );
}