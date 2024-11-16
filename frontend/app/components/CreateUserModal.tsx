"use client";

import { useState } from "react";
import { User } from "../lib/models/User";
import { AccessLevel, convertToStringAccessLevel } from "../lib/enums/AccessLevel";

type CreateUserModalProps = {
    isOpen: boolean,
    setIsOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export const CreateUserModal = ({isOpen, setIsOpen}: CreateUserModalProps) => {
    const emptyUser: User = {
        name: "",
        accessLevel: AccessLevel.Unknown,
        email: "",
        password: ""
    };

    const [formData, setFormData] = useState<User>(emptyUser); 

    const toggleModal = () => setIsOpen(!isOpen);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        setFormData((prev) => ({ ...prev, [name]: value }));
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        
        fetch("http://localhost:3410/protected/create_user", {
            method: "POST",
            body: JSON.stringify(formData)
        });

        toggleModal();
    };

    return (
        <div>
            {isOpen && (
                <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
                    <div className="w-full max-w-md p-6 bg-white rounded-lg shadow-lg">
                        <h2 className="text-2xl font-bold text-gray-800 mb-4">Fill in your details</h2>
                        <form onSubmit={handleSubmit}>
                            <div className="mb-4">
                                <label
                                    htmlFor="name"
                                    className="block text-sm font-medium text-gray-700"
                                >
                                    Name
                                </label>
                                <input
                                    type="text"
                                    id="name"
                                    name="name"
                                    value={formData.name}
                                    onChange={handleChange}
                                    className="w-full mt-1 p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                                    required
                                />
                            </div>

                            <div className="mb-4">
                                <label
                                    htmlFor="email"
                                    className="block text-sm font-medium text-gray-700"
                                >
                                    Email
                                </label>
                                <input
                                    type="email"
                                    id="email"
                                    name="email"
                                    value={formData.email}
                                    onChange={handleChange}
                                    className="w-full mt-1 p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                                    required
                                />
                            </div>

                            <div className="mb-4">
                                <label
                                    htmlFor="password"
                                    className="block text-sm font-medium text-gray-700"
                                >
                                    Password
                                </label>
                                <input
                                    type="password"
                                    id="password"
                                    name="password"
                                    value={formData.password}
                                    onChange={handleChange}
                                    className="w-full mt-1 p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                                    required
                                />
                            </div>

                            <div className="mb-4">
                                <label
                                    htmlFor="role"
                                    className="block text-sm font-medium text-gray-700"
                                >
                                    Role
                                </label>
                                <select
                                    id="role"
                                    name="role"
                                    value={formData.accessLevel && convertToStringAccessLevel(formData.accessLevel)}
                                    onChange={handleChange}
                                    className="w-full mt-1 p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500"
                                >
                                    <option value="Professor">Professor</option>
                                    <option value="Student">Student</option>
                                </select>
                            </div>

                            <div className="flex justify-end space-x-3">
                                <button
                                    type="button"
                                    onClick={toggleModal}
                                    className="px-4 py-2 bg-gray-300 rounded hover:bg-gray-400"
                                >
                                    Close
                                </button>
                                <button
                                    type="submit"
                                    className="px-4 py-2 text-white bg-blue-500 rounded hover:bg-blue-600"
                                    onClick={handleSubmit}
                                >
                                    Submit
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            )}
        </div>
    );
}

// import React, { useState } from "react";
// import { AddChoice } from "./AddChoice";
// import { User } from "../lib/models/User";
// import { AccessLevel } from "../lib/enums/AccessLevel";

// type CreateUserModalProps = {
//     hidden: boolean,
//     setIsHidden: React.Dispatch<React.SetStateAction<boolean>>;
// }

// export const CreateUserModal = ({hidden, setIsHidden}: CreateUserModalProps) => {
//     const [firstName, setFirstName] = useState("");
//     const [lastName, setLastName] = useState("");
//     const [email, setEmail] = useState("");
//     const [accessLevel, setAccessLevel] = useState("");
//     const [classes, setClasses] = useState<string[]>([]);
//     const [printers, setPrinters] = useState<string[]>([]);
//     const [printersCanAssign, setPrintersCanAssign] = useState(-1);

//     const convertToDbAccessLevel = (accessLevel: string): AccessLevel => {
//         switch (accessLevel) {
//             case "Admin":
//                 return AccessLevel.Admin;
//             case "Professor":
//                 return AccessLevel.Professor;
//             case "Student":
//                 return AccessLevel.Student;
//             default:
//                 return AccessLevel.Unknown
//         }
//     }

//     const onSubmit = () => {
//         const user: User = {
//             name: firstName + " " + lastName,
//             access_level: convertToDbAccessLevel(accessLevel),
//             email: email,
//             // classes: classes,
//             // printers: printers,
//             // printersCanAssign: printersCanAssign
//             password: "password"
//         };

//         fetch("http://localhost:3410/protected/create_user", {
//             method: "POST",
//             body: JSON.stringify(user)
//         });

//         setIsHidden(false);
//     }

//     return (
//         <div hidden={hidden} className={"h-full w-full flex flex-col"}>
//             <div>
//                 <label>First Name:</label>
//                 <input type="text" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
//             </div>
//             <div>
//                 <label>Last Name:</label>
//                 <input type="text" value={lastName} onChange={(e) => setLastName(e.target.value)} />
//             </div>
//             <div>
//                 <label>Email:</label>
//                 <input type="text" value={email} onChange={(e) => setEmail(e.target.value)} />
//             </div>
//             <div>
//                 <label>Professor or Student?</label>
//                 <input type="text" value={accessLevel} onChange={(e) => setAccessLevel(e.target.value)} />
//             </div>
//             <div>
//                 <label>Enter classes:</label>
//                 <AddChoice onChange={(classList: string[]) => setClasses(classList)} selections={classes} />
//             </div>
//             <div>
//                 <label>Enter printer ids:</label>
//                 <AddChoice onChange={(printerList: string[]) => setPrinters(printerList)} selections={printers} />
//             </div>
//             <div>
//                 <label># of printers allowed to manage</label>
//                 <input type="number" value={printersCanAssign} onChange={(e) => setPrintersCanAssign(e.target.valueAsNumber)} />
//             </div>

//             <button onClick={() => onSubmit()}>Submit</button>
//         </div>
//     );
// }