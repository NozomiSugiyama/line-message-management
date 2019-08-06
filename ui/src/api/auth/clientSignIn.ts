import { Token } from "src/api/type";
import config from "src/config";

export interface Credential {
    email: string;
    password: string;
}

export default async ({ email, password }: Credential, linkLine?: boolean) => {
    const response = await fetch(
        `${config.api.uri}/auth/client-sign-in${linkLine ? "?link-line=true" : ""}`,
        {
            method : "POST",
            body   : JSON.stringify({ email, password })
        }
    );

    if (!response.ok) {
        throw response;
    }

    return (await response.json()) as Token;
};
