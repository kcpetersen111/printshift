"use client";

import { useState } from "react";
import { AccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { CreateUserModal } from "./CreateUserModal";

type ProfessorPanelProps = {
    user: User
}

export const ProfessorPanel = ({user}: ProfessorPanelProps) => {
    const [openCreateUserModal, setOpenCreateUserModal] = useState(false);

    return(
        <div hidden={user.accessLevel !== AccessLevel.Professor} className="flex flex-col justify-start w-fit">
            <button onClick={() => setOpenCreateUserModal(true)} className="w-fit text-2xl bg-slate-300 rounded-md px-3 py-1.5 mt-20">Create Student</button>
            <CreateUserModal isOpen={openCreateUserModal} setIsOpen={setOpenCreateUserModal} access={user.accessLevel} />
        </div>
    );
}