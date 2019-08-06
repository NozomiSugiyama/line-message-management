import config from "src/config";

export default async (
    {
        id,
        message,
    }: {
        id: string,
        message: string
    }
) => {
    const response = await fetch(
        `${config.api.uri}/line_users/${id}/send_message`,
        {
            method: "POST",
            body   : JSON.stringify({ message }),
            headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
            }
        }
    );

    if (!response.ok) {
        throw response;
    }

    return;
};
