import { AccessLevel } from "../enums/AccessLevel"

export type User = {
    name: string,
    access_level: AccessLevel,
    email: string,
    // classes: string[],
    // printers: string[],
    // printersCanAssign: number
    password: string
}