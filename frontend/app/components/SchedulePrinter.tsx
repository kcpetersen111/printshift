import { useEffect, useState } from "react";
import { Printer } from "../lib/models/Printer";
import { Class } from "../lib/models/Class";

type SchedulePrinterProps = {
    isOpen: boolean,
    setIsOpen: React.Dispatch<React.SetStateAction<boolean>>,
    setTitle: React.Dispatch<React.SetStateAction<string>>
}

type SchedulePrinterRequest = {
    printerId: number,
    classId: number
}

export const SchedulePrinter = ({isOpen, setIsOpen, setTitle}: SchedulePrinterProps) => {
    const emptyScheduledPrinter: SchedulePrinterRequest = {
        printerId: -1,
        classId: -1
    }

    const [formData, setFormData] = useState<SchedulePrinterRequest>(emptyScheduledPrinter);
    const [printers, setPrinters] = useState<Printer[]>([]);
    const [classes, setClasses] = useState<Class[]>([]);

    const getAllPrinters = async () => {
        const response = fetch("http://localhost:3410/protected/list_printers", {
            method: "GET",
        });

        setPrinters(await ((await response).json() as Promise<Printer[]>))
    }
    const getAllClasses = async () => {
        const response = fetch("http://localhost:3410/protected/list_classes", {
            method: "GET",
        });

        setClasses(await ((await response).json() as Promise<Class[]>))
    }

    useEffect(() => {
        getAllPrinters();
        getAllClasses();
    }, [])

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>) => {
        const { name, value } = e.target;
        setFormData((prev) => ({ ...prev, [name]: Number.parseInt(value) }));
    };

    const toggleModal = () => setIsOpen(!isOpen);

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();

        setTitle(formData.classId.toString() + "-" + formData.printerId.toString());

        toggleModal();
    };

    return(
        <div>
            {isOpen && (
                <div className="fixed inset-0 flex items-center justify-center bg-black bg-opacity-50 z-50">
                    <div className="w-full max-w-md p-6 bg-white dark:bg-gray-800 rounded-lg shadow-lg">
                        <h2 className="text-2xl font-bold text-gray-800 dark:text-gray-100 mb-4">Fill in your details</h2>
                        <form onSubmit={handleSubmit}>

                            <div className="mb-4">
                                <label
                                    htmlFor="printerId"
                                    className="block text-sm font-medium text-gray-700 dark:text-gray-300"
                                >
                                    Printer
                                </label>
                                <select
                                    id="printerId"
                                    name="printerId"
                                    value={formData.printerId}
                                    onChange={handleChange}
                                    className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-700 text-gray-800 dark:text-gray-200 focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                >
                                    {printers.map((item, key) => (
                                        <option value={item.printerId} key={key}>{item.name}</option>
                                    ))}
                                </select>
                            </div>
                            <div className="mb-4">
                                <label
                                    htmlFor="classId"
                                    className="block text-sm font-medium text-gray-700 dark:text-gray-300"
                                >
                                    Class
                                </label>
                                <select
                                    id="classId"
                                    name="classId"
                                    value={formData.classId}
                                    onChange={handleChange}
                                    className="w-full mt-1 p-2 border border-gray-300 dark:border-gray-700 rounded-md bg-white dark:bg-gray-700 text-gray-800 dark:text-gray-200 focus:ring-blue-500 focus:border-blue-500 dark:focus:ring-blue-500 dark:focus:border-blue-500"
                                >
                                    {classes.map((item, key) => (
                                        <option value={item.id} key={key}>{item.name}</option>
                                    ))}
                                </select>
                            </div>

                            <div className="flex justify-end space-x-3">
                                <button
                                    type="button"
                                    onClick={toggleModal}
                                    className="px-4 py-2 bg-gray-300 dark:bg-gray-600 text-gray-800 dark:text-gray-200 rounded hover:bg-gray-400 dark:hover:bg-gray-500"
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