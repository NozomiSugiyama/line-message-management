
export interface Token {
    email: string;
    line_nonce: string;
}

export interface User {
    id: number;
    name: string;
    email: string;
}

export type Users = User[];

export interface LineUser {
    id: string;
    user_id: number;
    user: User;
    line_id: string;
    linked_account: "main";
    display_name?: string;
}

export type LineUsers = LineUser[];
