import { AccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";

type AdminPanelProps = {
    user: User
}

export const AdminPanel = ({user}: AdminPanelProps) => {
    return(
        <div hidden={user.accessLevel === AccessLevel.Admin}>
            <button>Create New User</button>
        </div>
    );
}