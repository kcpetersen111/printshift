"use client";

import { useEffect, useState } from "react";
import { Class } from "../lib/models/Class";
import { Professor } from "../lib/models/Professor";

type CreateClassModalProps = {
    isOpen: boolean,
    setIsOpen: React.Dispatch<React.SetStateAction<boolean>>
}

export const CreateClassModal = ({isOpen, setIsOpen}: CreateClassModalProps) => {
    const emptyClass: Class = {
        name: "",
        description: "",
        professorId: -1,
        isActive: true,
        id: -1
    };

    const [formData, setFormData] = useState<Class>(emptyClass);
    const [professors, setProfessors] = useState<Professor[]>([]);

    const getAllProfessors = async () => {
        const response = fetch("http://localhost:3410/protected/list_professors", {
            method: "GET",
        });

        setProfessors(await ((await response).json() as Promise<Professor[]>))
    }

    useEffect(() => {
        getAllProfessors();
    }, [isOpen])



    const toggleModal = () => setIsOpen(!isOpen);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        if (name == "professorId") {
            setFormData((prev) => ({ ...prev, [name]: Number.parseInt(value) }));
        } else {
            setFormData((prev) => ({ ...prev, [name]: value }));
        }
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        
        fetch("http://localhost:3410/protected/class", {
            method: "POST",
            body: JSON.stringify(formData)
        });

        toggleModal();
    };

    return (
        <div>
            {isOpen && (
                <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
                    <div className="w-full max-w-md p-6 bg-white dark:bg-gray-800 rounded-lg shadow-lg">
                        <h2 className="text-2xl font-bold text-gray-800 dark:text-white mb-4">Fill in your details</h2>
                        <form onSubmit={handleSubmit}>
                            <div className="mb-4">
                                <label
                                    htmlFor="name"
                                    className="block text-sm font-medium text-gray-700 dark:text-gray-300"
                                >
                                    Name
                                </label>
                                <input
                                    type="text"
                                    id="name"
                                    name="name"
                                    value={formData.name}
                                    onChange={handleChange}
                                    className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-blue-500 dark:focus:border-blue-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                                    required
                                />
                            </div>

                            <div className="mb-4">
                                <label
                                    htmlFor="description"
                                    className="block text-sm font-medium text-gray-700 dark:text-gray-300"
                                >
                                    Description
                                </label>
                                <input
                                    type="description"
                                    id="description"
                                    name="description"
                                    value={formData.description}
                                    onChange={handleChange}
                                    className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-600 rounded-md focus:ring-blue-500 dark:focus:ring-blue-400 focus:border-blue-500 dark:focus:border-blue-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                                    required
                                />
                            </div>
                            <div className="mb-4">
                                <label
                                    htmlFor="professorId"
                                    className="block text-sm font-medium text-gray-700 dark:text-gray-300"
                                >
                                    ProfessorId
                                </label>
                                <select
                                    id="professorId"
                                    name="professorId"
                                    value={formData.professorId}
                                    onChange={handleChange}
                                    className="w-full mt-1 p-2 border border-gray-300 rounded-md focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-400 focus:border-blue-500 dark:focus:border-blue-400 bg-white dark:bg-gray-700 text-gray-900 dark:text-gray-100"
                                >
                                    {professors.map((item, key) => (
                                        <option value={item.professorId} key={key}>{item.name}</option>
                                    ))}
                                </select>
                            </div>

                            <div className="flex justify-end space-x-3">
                                <button
                                    type="button"
                                    onClick={toggleModal}
                                    className="px-4 py-2 bg-gray-300 dark:bg-gray-600 text-white rounded hover:bg-gray-400 dark:hover:bg-gray-500"
                                >
                                    Close
                                </button>
                                <button
                                    type="submit"
                                    className="px-4 py-2 text-white bg-blue-500 rounded hover:bg-blue-600 dark:bg-blue-600 dark:hover:bg-blue-500"
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

