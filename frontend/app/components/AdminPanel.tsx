"use client";

import { useState } from "react";
import { AccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { CreateUserModal } from "./CreateUserModal";
import { CreateClassModal } from "./CreateClassModal";

type AdminPanelProps = {
    user: User
}

export const AdminPanel = ({user}: AdminPanelProps) => {
    const [openCreateUserModal, setOpenCreateUserModal] = useState(false);
    const [openCreateClassModal, setOpenCreateClassModal] = useState(false);

    return(
        <div hidden={user.accessLevel !== AccessLevel.Admin} className="flex flex-col justify-start w-fit">
            <button onClick={() => setOpenCreateUserModal(true)} className="w-fit text-2xl bg-slate-300 rounded-md px-3 py-1.5 mt-20">Create New User</button>
            <CreateUserModal isOpen={openCreateUserModal} setIsOpen={setOpenCreateUserModal} access={AccessLevel.Admin}/>
            <button onClick={() => setOpenCreateClassModal(true)} className="w-fit text-2xl bg-slate-300 rounded-md px-3 py-1.5 my-32">Create New Class</button>
            <CreateClassModal isOpen={openCreateClassModal} setIsOpen={setOpenCreateClassModal} />
            <button className="w-fit text-2xl bg-slate-300 rounded-md px-3 py-1.5">Manage Printers</button>
        </div>
    );
}