import { Person } from "./person";

export interface Star {
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt?: any;
    Comment: string;
    Recipient: Person;
    RecipientID: number;
    Sender: Person;
    SenderID: number;
}
