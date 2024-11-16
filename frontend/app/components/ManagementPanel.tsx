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
                throw new Error("Access Level unrecognized: " + user.access_level);            
        }
    }

    return (
        <div className="w-fit">
            <h1 className="text-8xl mt-10">{convertToStringAccessLevel(user.access_level)} Panel</h1>
            <AdminPanel user={user} />
        </div>
    );
}