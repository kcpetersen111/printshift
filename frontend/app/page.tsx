"use client";
import { useState } from "react";
import { ManagementPanel } from "./components/ManagementPanel";
import { AccessLevel } from "./lib/enums/AccessLevel";
import { User } from "./lib/models/User";
import { LoginRequest } from "./lib/models/LoginRequest";
import { CreateUserResponse } from "./lib/models/CreateUserResponse";

export default function Home() {
    const [currentUser, setCurrentUser] = useState<User | null>(null);

    const login = async (email: string, password: string) => {
        const loginRequest: LoginRequest = {
            email,
			password
        };

		const response = fetch("http://localhost:3410/login", {
            method: "POST",
            body: JSON.stringify(loginRequest)
        }).json() as Promise<CreateUserResponse>;

		if ((await response)) {
			setCurrentUser((await response).user);
		} else {
			throw new Error("Unknown error: User not authenticated.");
		}
		
    };
    
    // Creates a test user
    
	// const user: User = {
	// 	name: "Admin Dummy User",
	// 	access_level: AccessLevel.Admin,
	// 	email: "dummy@user.com",
	// 	// classes: [],
	// 	// printers: [],
	// 	// printersCanAssign: 1000
	// 	password: "pass"
	// };
    
    //For the admin panel

	// return (
	// 	<div className="flex flex-col items-center">
	// 		<ManagementPanel user={user} />
	// 	</div>
	// );

    return (
        <div className="flex flex-col items-center">
            <h1 className="text-4xl mb-6">Login</h1>

            {/* User Creation Form */}
            <form
                onSubmit={(e) => {
                    e.preventDefault();
                    const form = e.target as HTMLFormElement;
                    const email = (form.elements.namedItem("email") as HTMLInputElement).value;
					const password = (form.elements.namedItem("password") as HTMLInputElement).value;
                    login(email, password);
                    form.reset();
                }}
                className="mb-6"
            >
                <div className="mb-4">
                    <label className="block text-lg">Email:</label>
                    <input type="email" name="email" required className="border p-2" />
                </div>
                <div className="mb-4">
                    <label className="block text-lg">Password:</label>
                    <input type="password" name="password" required className="border p-2" />
                </div>
                <button type="submit" className="bg-blue-500 text-white p-2 rounded">
                    Login
                </button>
            </form>

            {/* Management Panel for Current User */}
            {currentUser && (
                <ManagementPanel user={currentUser} />
            )}
        </div>
    );
}
