import React, { useState, Fragment } from "react";
import NotificationComponent from "src/components/atoms/Notification";
import NotificationContext from "src/contexts/NotificationContext";

interface Notification {
    type: "info" | "error";
    message: string;
    key: number;
    close: () => void;
}

export default (
    {
        children
    }: {
        children?: React.ReactNode
    }
) => {

    const [notifications, setNotification] = useState<Notification[]>([]);

    return (
        <Fragment>
            <NotificationContext.Provider
                value={{
                    ErrorComponent: ({ message }) => <NotificationComponent type="error" message={message}/>,
                    notification: async (type: "info" | "error", message: string) => new Promise((resolve) => {
                        const key = Date.now();
                        setNotification(
                            notifications.concat({
                                type,
                                message,
                                key,
                                close: () => {
                                    setNotification(
                                        notifications.filter(y => key !== y.key)
                                    );
                                    resolve();
                                }
                            })
                        );
                    })
                }}
            >
                {children}
            </NotificationContext.Provider>
            {notifications.map(x =>
                <NotificationComponent
                    type={x.type}
                    key={x.key}
                    message={x.message}
                    open={true}
                    onClose={x.close}
                />
            )}
        </Fragment>
    );
};
