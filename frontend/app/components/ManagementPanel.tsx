import { convertToStringAccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { AdminPanel } from "./AdminPanel";
import { ProfessorPanel } from "./ProfessorPanel";

type ManagementPanelProps = {
    user: User
}

export const ManagementPanel = ({user}: ManagementPanelProps) => {

    return (
        <div className="w-8/12">
            <h1 className="text-6xl mt-10 text-center">{convertToStringAccessLevel(user.accessLevel)} Panel</h1>
            <AdminPanel user={user} />
            <ProfessorPanel user={user} />
        </div>
    );
}
