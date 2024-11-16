"use client";

import { useState } from "react";
import { AccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { CreateUserModal } from "./CreateUserModal";
import { CreateClassModal } from "./CreateClassModal";

type AdminPanelProps = {
    user: User;
};

export const AdminPanel = ({ user }: AdminPanelProps) => {
    const [openCreateUserModal, setOpenCreateUserModal] = useState(false);
    const [openCreateClassModal, setOpenCreateClassModal] = useState(false);

    return (
        <div hidden={user.accessLevel !== AccessLevel.Admin} className="flex flex-col justify-start w-full mt-10">
            <div className="flex flex-col space-y-4 w-full">
                {/* Row 1: Create New User */}
                <div className="bg-gray-200 dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-600 transition p-4 rounded-md w-full">
                    <button
                        onClick={() => setOpenCreateUserModal(true)}
                        className="flex justify-between items-center w-full text-2xl text-foreground px-4"
                    >
                        <span className="text-left">Create New User</span>
                        <span className="text-right">&gt;</span>
                    </button>
                    <CreateUserModal isOpen={openCreateUserModal} setIsOpen={setOpenCreateUserModal} />
                </div>

                {/* Row 2: Manage Printers */}
                <div className="bg-gray-100 dark:bg-gray-800 hover:bg-gray-200 dark:hover:bg-gray-700 transition p-4 rounded-md w-full">
                    <button className="flex justify-between items-center w-full text-2xl text-foreground px-4">
                        <span className="text-left">Manage Printers</span>
                        <span className="text-right">&gt;</span>
                    </button>
                </div>

                {/* Row 3: Create New Class */}
                <div className="bg-gray-200 dark:bg-gray-700 hover:bg-gray-300 dark:hover:bg-gray-600 transition p-4 rounded-md w-full">
                    <button
                        onClick={() => setOpenCreateClassModal(true)}
                        className="flex justify-between items-center w-full text-2xl text-foreground px-4"
                    >
                        <span className="text-left">Create New Class</span>
                        <span className="text-right">&gt;</span>
                    </button>
                    <CreateClassModal isOpen={openCreateClassModal} setIsOpen={setOpenCreateClassModal} />
                </div>
            </div>
        </div>
    );
};

