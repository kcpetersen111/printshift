"use client";

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
        <div hidden={user.access_level !== AccessLevel.Admin} className="flex flex-col justify-start w-fit">
            <button onClick={() => {setCreateUserModal(true)}} className="w-fit text-2xl bg-slate-300 rounded-md px-3 py-1.5 mt-20">Create New User</button>
            <CreateUserModal hidden={!createUserModal} setIsHidden={setCreateUserModal} />
            <button className="w-fit text-2xl bg-slate-300 rounded-md px-3 py-1.5 my-32">Manage Users</button>
            <button className="w-fit text-2xl bg-slate-300 rounded-md px-3 py-1.5">Manage Printers</button>
        </div>
    );
}