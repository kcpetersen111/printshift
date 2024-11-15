import { useState } from "react";
import { AccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { CreateUserModal } from "./CreateUserModal";

type AdminPanelProps = {
    user: User
}

export const AdminPanel = ({user}: AdminPanelProps) => {
    const [createUserModal, setCreateUserModal] = useState(false);

    return(
        <div hidden={user.accessLevel === AccessLevel.Admin}>
            <button onClick={() => setCreateUserModal(true)}>Create New User</button>
            <CreateUserModal hidden={!createUserModal} />
            <button>Manage Users</button>
            <button>Manage Printers</button>
        </div>
    );
}