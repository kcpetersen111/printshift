"use client";

import { useState } from "react";
import { AccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { CreateUserModal } from "./CreateUserModal";
import { CreatePrinterModal } from "./CreatePrinterModal";

type ProfessorPanelProps = {
    user: User
}

export const ProfessorPanel = ({user}: ProfessorPanelProps) => {
    const [openCreateUserModal, setOpenCreateUserModal] = useState(false);
    const [openCreatePrinterModal, setOpenCreatePrinterModal] = useState(false);

    return(
        <div hidden={user.accessLevel !== AccessLevel.Professor} className="flex flex-col justify-start w-full mt-10">
            <div className="flex flex-col space-y-4 w-full">
                <div className="bg-gray-200 dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-600 transition p-4 rounded-md w-full">
                    <button onClick={() => setOpenCreateUserModal(true)} className="flex justify-between items-center w-full text-2xl text-foreground px-4">
                        <span className="text-left">Create Student</span>
                        <span className="text-right">&gt;</span>
                    </button>
                    <CreateUserModal isOpen={openCreateUserModal} setIsOpen={setOpenCreateUserModal} access={user.accessLevel} />
                </div>
                <div className="bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700 transition p-4 rounded-md w-full">
                    <button onClick={() => setOpenCreatePrinterModal(true)} className="flex justify-between items-center w-full text-2xl text-foreground px-4">
                        <span className="text-left">Create Printers</span>
                        <span className="text-right">&gt;</span>
                    </button>
                    <CreatePrinterModal isOpen={openCreatePrinterModal} setIsOpen={setOpenCreatePrinterModal} />
                </div>
            </div>
        </div>
    );
}