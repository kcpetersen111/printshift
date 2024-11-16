import { convertToStringAccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { AdminPanel } from "./AdminPanel";

type ManagementPanelProps = {
    user: User
}

export const ManagementPanel = ({user}: ManagementPanelProps) => {

    return (
        <div className="w-fit">
            <h1 className="text-8xl mt-10">{convertToStringAccessLevel(user.accessLevel)} Panel</h1>
            <AdminPanel user={user} />
        </div>
    );
}
