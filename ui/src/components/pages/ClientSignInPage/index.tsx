import { Button, FormControl, IconButton, Input, InputAdornment, InputLabel, TextField } from "@material-ui/core";
import VisibilityIcon from "@material-ui/icons/Visibility";
import VisibilityOffIcon from "@material-ui/icons/VisibilityOff";
import React, { useCallback, useContext, useState } from "react";
import { RouteChildrenProps } from "react-router";
import clientSignIn from "src/api/auth/clientSignIn";
import toObjectFromURIQuery from "src/api/toObjectFromURIQuery";
import Page from "src/components/atoms/Page";
import Header from "src/components/molecules/Header";
import NotificationContext from "src/contexts/NotificationContext";
import styled from "styled-components";

export type TopPageProps = React.ComponentProps<typeof Page> & RouteChildrenProps<{problemId: string}>;

export default (props: TopPageProps) => {
    const [passwordIsVisible, setPasswordVisibility] = useState<boolean>(false);

    const notificationContext = useContext(NotificationContext);

    const linkLineToken = (toObjectFromURIQuery(location.search) || { "link-line-token": null })["link-line-token"];
    const submitForm = useCallback(
        async (e: React.FormEvent<HTMLFormElement>) => {
            e.preventDefault();

            const email = (e.target as any).elements["sign-in-email"].value;
            const password = (e.target as any).elements["sign-in-password"].value;

            const token = await clientSignIn({ email, password }, !!linkLineToken);
            console.log(token);

            location.href = `https://access.line.me/dialog/bot/accountLink?linkToken=${linkLineToken}&nonce=${token.line_nonce}`;
        },
        [notificationContext.notification]
    );

    return (
        <Page
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
                <div>Link line token: {linkLineToken}</div>
            </div>
        </Page>
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
