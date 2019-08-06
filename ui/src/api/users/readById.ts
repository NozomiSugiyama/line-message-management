import { User } from "src/api/type";
import config from "src/config";

export default async ({ id }: { id: string }) => {
    const response = await fetch(
        `${config.api.uri}/users/${id}`
    );

    if (!response.ok) {
        throw response;
    }

    return (await response.json()) as User;
};
