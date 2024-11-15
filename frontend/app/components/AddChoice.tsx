type AddChoiceProps = {
    onChange: (selections: string[]) => void,
    selections: string[]
}

export const AddChoice = ({onChange, selections}: AddChoiceProps) => {

    return (
        <div>
            <input type="text" onChange={(e) => onChange([...selections, e.target.value])}/>
            <ul>
                {selections.map((c, key) => (
                    <li key={key}>{c}</li>
                ))}
            </ul>
        </div>
    );
}