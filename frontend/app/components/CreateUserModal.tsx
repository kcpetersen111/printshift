"use client";

import React, { useState } from "react";
import { AddChoice } from "./AddChoice";
import { User } from "../lib/models/User";
import { AccessLevel } from "../lib/enums/AccessLevel";

type CreateUserModalProps = {
    hidden: boolean,
    setIsHidden: React.Dispatch<React.SetStateAction<boolean>>;
}

export const CreateUserModal = ({hidden, setIsHidden}: CreateUserModalProps) => {
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [email, setEmail] = useState("");
    const [accessLevel, setAccessLevel] = useState("");
    const [classes, setClasses] = useState<string[]>([]);
    const [printers, setPrinters] = useState<string[]>([]);
    const [printersCanAssign, setPrintersCanAssign] = useState(-1);

    const convertToDbAccessLevel = (accessLevel: string): AccessLevel => {
        switch (accessLevel) {
            case "Admin":
                return AccessLevel.Admin;
            case "Professor":
                return AccessLevel.Professor;
            case "Student":
                return AccessLevel.Student;
            default:
                return AccessLevel.Unknown
        }
    }

    const onSubmit = () => {
        const user: User = {
            name: firstName + " " + lastName,
            accessLevel: convertToDbAccessLevel(accessLevel),
            email: email,
            classes: classes,
            printers: printers,
            printersCanAssign: printersCanAssign
        };

        // TODO: API Call CreateUser
        setIsHidden(false);
    }

    return (
        <div hidden={hidden} className={"h-full w-full flex flex-col"}>
            <div>
                <label>First Name:</label>
                <input type="text" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
            </div>
            <div>
                <label>Last Name:</label>
                <input type="text" value={lastName} onChange={(e) => setLastName(e.target.value)} />
            </div>
            <div>
                <label>Email:</label>
                <input type="text" value={email} onChange={(e) => setEmail(e.target.value)} />
            </div>
            <div>
                <label>Professor or Student?</label>
                <input type="text" value={accessLevel} onChange={(e) => setAccessLevel(e.target.value)} />
            </div>
            <div>
                <label>Enter classes:</label>
                <AddChoice onChange={(classList: string[]) => setClasses(classList)} selections={classes} />
            </div>
            <div>
                <label>Enter printer ids:</label>
                <AddChoice onChange={(printerList: string[]) => setPrinters(printerList)} selections={printers} />
            </div>
            <div>
                <label># of printers allowed to manage</label>
                <input type="number" value={printersCanAssign} onChange={(e) => setPrintersCanAssign(e.target.valueAsNumber)} />
            </div>

            <button onClick={() => onSubmit()}>Submit</button>
        </div>
    );
}