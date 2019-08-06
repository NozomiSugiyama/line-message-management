import config from "src/config";

export default async (
    {
        id
    }: {
        id: string
    }
) => {
    const response = await fetch(
        `${config.api.uri}/line_users/${id}/send_test_messages`,
        {
            method: "POST"
        }
    );

    if (!response.ok) {
        throw response;
    }

    return;
};
