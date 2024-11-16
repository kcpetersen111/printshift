import { convertToStringAccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { AdminPanel } from "./AdminPanel";

type ManagementPanelProps = {
    user: User
}

export const ManagementPanel = ({ user }: ManagementPanelProps) => {

    return (
        <div className="w-[70%] flex flex-col items-center space-y-6 text-gray-800 dark:text-white">
            <h1 className="text-6xl font-semibold mt-12 text-left">{convertToStringAccessLevel(user.accessLevel)} Panel</h1>
            <AdminPanel user={user} />
        </div>
    );
}

