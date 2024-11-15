import { AccessLevel } from "../enums/AccessLevel"

export type User = {
    name: string,
    accessLevel: AccessLevel,
    email: string,
    classes: string[],
    printers: string[],
    printersCanAssign: number
}