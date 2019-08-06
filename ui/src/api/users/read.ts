import { Users } from "src/api/type";
import config from "src/config";

export default async () => {
    const response = await fetch(
        `${config.api.uri}/users`
    );

    if (!response.ok) {
        throw response;
    }

    return (await response.json()) as Users;
};
