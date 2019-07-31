import React from "react";
import { RouteChildrenProps } from "react-router";
import Page from "src/components/atoms/Page";
import Header from "src/components/molecules/Header";
import NotFound from "src/components/molecules/NotFound";
import Content from "src/components/pages/NotFoundPage/Content";

export type NotFoundPageProps = React.ComponentProps<typeof Page> & RouteChildrenProps<{problemId: string}>;

export default (props: NotFoundPageProps) => (
    <Page
        {...props}
    >
        <Header
            appTitle="Not Found"
        />
        <Content>
            <NotFound/>
        </Content>
    </Page>
);
