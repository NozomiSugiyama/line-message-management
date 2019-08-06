import { LineUsers } from "src/api/type";
import config from "src/config";

export default async () => {
    const response = await fetch(
        `${config.api.uri}/line_users`
    );

    if (!response.ok) {
        throw response;
    }

    return (await response.json()) as LineUsers;
};
