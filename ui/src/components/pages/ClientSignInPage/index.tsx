import { Button, FormControl, IconButton, Input, InputAdornment, InputLabel, TextField } from "@material-ui/core";
import VisibilityIcon from "@material-ui/icons/Visibility"
import VisibilityOffIcon from "@material-ui/icons/VisibilityOff"
import cryptoRandomString from "crypto-random-string";
import React, { useCallback, useContext, useState } from "react";
import { RouteChildrenProps } from "react-router";
import toObjectFromURIQuery from "src/api/toObjectFromURIQuery";
import Header from "src/components/molecules/Header";
import Host from "src/components/pages/ClientSignInPage/Host";
import NotificationContext from "src/contexts/NotificationContext";
import styled from "styled-components";

export type TopPageProps = React.ComponentProps<typeof Host> & RouteChildrenProps<{problemId: string}>;

export default (props: TopPageProps) => {
    const [passwordIsVisible, setPasswordVisibility] = useState<boolean>(false);

    const notificationContext = useContext(NotificationContext);

    const linkToken = (toObjectFromURIQuery(location.search) || { "link-token": null })["link-token"];
    const submitForm = useCallback(
        async (e: React.FormEvent<HTMLFormElement>) => {
            e.preventDefault();

            const email = (e.target as any).elements["sign-in-email"].value;
            const password = (e.target as any).elements["sign-in-password"].value;

            const nonce = cryptoRandomString({ length: 255, type: "base64" });
            await notificationContext.notification("info", `${email}:${password}`);
            await notificationContext.notification("info", `nonce: ${nonce}`);
            location.href = `https://access.line.me/dialog/bot/accountLink?linkToken=${linkToken}&nonce=${nonce}`;
        },
        [notificationContext.notification]
    );

    return (
        <Host
            {...props}
        >
            <Header
                appTitle="Client Sign in"
            />
            <form
                onSubmit={submitForm}
            >
                <FormContent>
                    <TextField
                        id="sign-in-email"
                        label="Mail address"
                        margin="normal"
                        type="email"
                        required
                    />
                    <FormControl
                        margin="normal"
                    >
                        <InputLabel htmlFor="sign-in-password">Password</InputLabel>
                        <Input
                            id="sign-in-password"
                            type={passwordIsVisible ? "text" : "password"}
                            required
                            endAdornment={
                                <InputAdornment position="end">
                                    <IconButton
                                        onClick={() => setPasswordVisibility(x => !x)}
                                    >
                                        {passwordIsVisible ? <VisibilityIcon /> : <VisibilityOffIcon />}
                                    </IconButton>
                                </InputAdornment>
                            }
                        />
                    </FormControl>
                </FormContent>
                <FormActions>
                    <Button
                        component="button"
                        color="primary"
                        type="submit"
                    >
                        Sign in
                    </Button>
                </FormActions>
            </form>
            <div>
                <div>Link token: {linkToken}</div>
            </div>
        </Host>
    );
};

const FormContent = styled.div`
    display: flex;
    flex-direction: column;
`;

const FormActions = styled.div`
    display: flex;
    justify-content: flex-end;
`;
