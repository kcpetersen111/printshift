"use client";

import { useState } from "react"

type AddChoiceProps = {
    onChange: (selections: string[]) => void,
    selections: string[]
}

export const AddChoice = ({onChange, selections}: AddChoiceProps) => {
    const [choice, setChoice] = useState("");
    const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
        if (event.key === "Enter") {
            if (choice !== "") {
                onChange([...selections, choice])
                setChoice("");
            }
        }
    }

    return (
        <div>
            <input type="text" value={choice} onChange={(e) => setChoice(e.target.value)} onKeyDown={handleKeyDown}/>
            <ul>
                {selections.map((c, key) => (
                    <li key={key}>{c}</li>
                ))}
            </ul>
        </div>
    );
}