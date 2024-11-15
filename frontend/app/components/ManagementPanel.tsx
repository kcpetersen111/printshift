import { AccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { AdminPanel } from "./AdminPanel";

type ManagementPanelProps = {
    user: User
}

export const ManagementPanel = ({user}: ManagementPanelProps) => {
    const convertToStringAccessLevel = (accessLevel: AccessLevel) => {
        switch (accessLevel) {
            case AccessLevel.Admin:
                return "Admin";
            case AccessLevel.Professor:
                return "Professor";
            case AccessLevel.Student:
                return "Student";
            default:
                throw new Error("Access Level unrecognized: " + user.accessLevel);            
        }
    }

    return (
        <div>
            <h1>{convertToStringAccessLevel(user.accessLevel)} Panel</h1>
            <AdminPanel user={user} />
        </div>
    );
}