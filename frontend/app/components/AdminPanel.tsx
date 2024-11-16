"use client";

import { useState } from "react";
import { AccessLevel } from "../lib/enums/AccessLevel";
import { User } from "../lib/models/User";
import { CreateUserModal } from "./CreateUserModal";
import { CreateClassModal } from "./CreateClassModal";

type AdminPanelProps = {
    user: User
}

export const AdminPanel = ({ user }: AdminPanelProps) => {
    const [openCreateUserModal, setOpenCreateUserModal] = useState(false);
    const [openCreateClassModal, setOpenCreateClassModal] = useState(false);

    return (
        <div hidden={user.accessLevel !== AccessLevel.Admin} className="flex flex-col justify-start w-full max-w-4xl mx-auto space-y-4">
            {/* Create New User Button */}
            <div
                onClick={() => setOpenCreateUserModal(true)}
                className="cursor-pointer flex justify-start w-full py-2 px-6 bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600 rounded-lg shadow-md transition-all"
            >
                <button className="text-left w-full text-xl text-black dark:text-white rounded-md px-4 py-2">Create New User</button>
            </div>

            <CreateUserModal isOpen={openCreateUserModal} setIsOpen={setOpenCreateUserModal} />
            
            {/* Create New Class Button */}
            <div
                onClick={() => setOpenCreateClassModal(true)}
                className="cursor-pointer flex justify-start w-full py-2 px-6 bg-gray-200 hover:bg-gray-300 dark:bg-gray-800 dark:hover:bg-gray-700 rounded-lg shadow-md transition-all"
            >
                <button className="text-left w-full text-xl text-black dark:text-white rounded-md px-4 py-2">Create New Class</button>
            </div>
            
            <CreateClassModal isOpen={openCreateClassModal} setIsOpen={setOpenCreateClassModal} />
            
            {/* Manage Printers Button */}
            <div className="cursor-pointer flex justify-start w-full py-2 px-6 bg-gray-100 hover:bg-gray-200 dark:bg-gray-700 dark:hover:bg-gray-600 rounded-lg shadow-md transition-all">
                <button className="text-left w-full text-xl text-black dark:text-white rounded-md px-4 py-2">Manage Printers</button>
            </div>
        </div>
    );
}

