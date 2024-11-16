"use client";
import { useState } from "react";
import { ManagementPanel } from "./components/ManagementPanel";
import { AccessLevel } from "./lib/enums/AccessLevel";
import { User } from "./lib/models/User";

export default function Home() {
    const [users, setUsers] = useState<User[]>([]);
    const [currentUser, setCurrentUser] = useState<User | null>(null);

    const createUser = (name: string, email: string, accessLevel: AccessLevel) => {
        const newUser: User = {
            name,
            email,
            accessLevel,
            classes: [],
            printers: [],
            printersCanAssign: 0, // Default value or adjust as needed
        };
        setUsers([...users, newUser]);
    };
    
    // Creates a test user
    
	//const user: User = {
	//	name: "Admin Dummy User",
	//	accessLevel: AccessLevel.Admin,
	//	email: "dummy@user.com",
	//	classes: [],
	//	printers: [],
	//	printersCanAssign: 1000
	//};
    
    //For the admin panel

	//return (
	//	<div className="flex flex-col items-center">
	//		<ManagementPanel user={user} />
	//	</div>
	//);

    return (
        <div className="flex flex-col items-center">
            <h1 className="text-4xl mb-6">User Management</h1>

            {/* User Creation Form */}
            <form
                onSubmit={(e) => {
                    e.preventDefault();
                    const form = e.target as HTMLFormElement;
                    const name = (form.elements.namedItem("name") as HTMLInputElement).value;
                    const email = (form.elements.namedItem("email") as HTMLInputElement).value;
                    const accessLevel = parseInt(
                        (form.elements.namedItem("accessLevel") as HTMLSelectElement).value
                    ) as AccessLevel;
                    createUser(name, email, accessLevel);
                    form.reset();
                }}
                className="mb-6"
            >
                <div className="mb-4">
                    <label className="block text-lg">Name:</label>
                    <input type="text" name="name" required className="border p-2" />
                </div>
                <div className="mb-4">
                    <label className="block text-lg">Email:</label>
                    <input type="email" name="email" required className="border p-2" />
                </div>
                <div className="mb-4">
                    <label className="block text-lg">Access Level:</label>
                    <select name="accessLevel" className="border p-2" required>
                        <option value={AccessLevel.Admin}>Admin</option>
                        <option value={AccessLevel.Professor}>Professor</option>
                        <option value={AccessLevel.Student}>Student</option>
                    </select>
                </div>
                <button type="submit" className="bg-blue-500 text-white p-2 rounded">
                    Create User
                </button>
            </form>

            {/* Display User List */}
            <div className="mb-6">
                <h2 className="text-2xl mb-4">User List</h2>
                <ul>
                    {users.map((user, index) => (
                        <li key={index} className="mb-2">
                            <button
                                onClick={() => setCurrentUser(user)}
                                className="text-blue-500 underline"
                            >
                                {user.name} ({user.email})
                            </button>
                        </li>
                    ))}
                </ul>
            </div>

            {/* Management Panel for Current User */}
            {currentUser && (
                <ManagementPanel user={currentUser} />
            )}
        </div>
    );
}
