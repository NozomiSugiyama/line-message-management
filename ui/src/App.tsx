import React from "react";
import { BrowserRouter, Route, Switch } from "react-router-dom";
import Root from "src/Root";
import {
    ClientSignInPage,
    LineUsersPage,
    NotFoundPage,
    TopPage
} from "src/Routes";

export default () => (
    <BrowserRouter>
        <Root>
            <Switch>
                <Route
                    path="/"
                    component={TopPage}
                    exact
                />
                <Route
                    path="/line-users"
                    component={LineUsersPage}
                    exact
                />
                <Route
                    path="/top"
                    component={TopPage}
                    exact
                />
                <Route
                    path="/client-sign-in"
                    component={ClientSignInPage}
                    exact
                />
                <Route component={NotFoundPage}/>
            </Switch>
        </Root>
    </BrowserRouter>
);
