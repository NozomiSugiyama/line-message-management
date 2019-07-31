import { Drawer } from "@material-ui/core";
import React, { useState } from "react";
import Navigator from "src/components/molecules/Navigator";
import DrawerContext from "src/contexts/DrawerContext";
import styled from "styled-components";

export default (
    {
        children
    }: {
        children?: React.ReactNode
    }
) => {
    const [drawerOpened, setDrawerOpen] = useState<boolean>(false);

    const toggleDrawer = () => setDrawerOpen(!drawerOpened);
    const hideDrawer = () => drawerOpened && setDrawerOpen(false);
    const showDrawer = () => !drawerOpened && setDrawerOpen(true);

    return (
        <Host>
            <div>
                <Drawer
                    variant="temporary"
                    anchor={"left"}
                    open={drawerOpened}
                    onClose={toggleDrawer}
                    ModalProps={{ keepMounted: true }}
                >
                    <Navigator/>
                </Drawer>
            </div>
            <div>
                <Drawer
                    variant="permanent"
                    open
                >
                    <Navigator/>
                </Drawer>
            </div>
            <Content>
                <DrawerContext.Provider
                    value={{
                        toggle: toggleDrawer,
                        hide: hideDrawer,
                        show: showDrawer
                    }}
                >
                    {children}
                </DrawerContext.Provider>
            </Content>
        </Host>
    );
};

const Host = styled.div`
    background-color: #fafbfd;
    > :nth-child(1) {
        display: none;
    }
    > :nth-child(2) {
        display: flex;
    }

    @media (max-width: 767px) {
        > :nth-child(1) {
            display: flex;
        }
        > :nth-child(2) {
            display: none;
        }
    }
`;

const Content = styled.main`
    position: relative;
    width: calc(100% - 17rem);
    margin-left: 17rem;
    @media (max-width: 767px) {
        width: 100%;
        margin-left: 0rem;
    }
`;
