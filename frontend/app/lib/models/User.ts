import { AccessLevel } from "../enums/AccessLevel"

export type User = {
    name: string,
    accessLevel: AccessLevel,
    email: string,
    password: string
}