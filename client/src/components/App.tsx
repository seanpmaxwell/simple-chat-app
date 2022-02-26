import React, { useCallback, useEffect } from 'react';
import { BrowserRouter, Routes, Route } from "react-router-dom";

import CssBaseline from '@mui/material/CssBaseline';

import Home from './Home/Home';
import NoPage from './NoPage/NoPage';
import './App.css';
import TopBar from './TopBar/Topbar';
import Users from './Users/Users';
import Chat from './Chat/Chat';
import authHttp, { ISessionData } from './shared/http/auth-http';
import { useSetState } from './shared/hooks';


interface IState {
    sessionData: ISessionData;
}

function App() {
    const [state, setState] = useSetState(init());
    // Set fetch-session-data function
    const fetchSessionData = useCallback(() => {
        return getSessionData().then(sessionData => setState({sessionData}));
    }, [setState])
    // On load
    useEffect(() => {
        fetchSessionData()
    }, [fetchSessionData]);
    // Return
    return (
        <div>
            <React.Fragment>
                <CssBaseline />
                <BrowserRouter>
                    <Routes>
                        <Route
                            path="/"
                            element={
                                <TopBar
                                    sessionData={state.sessionData}
                                    fetchSessionData={() => fetchSessionData()}
                                />
                            }
                        >
                            <Route
                                index={true}
                                element={
                                    <Home fetchSessionData={() => fetchSessionData()}/>
                                }
                            />
                            <Route path="users" element={<Users />} />
                            <Route path="chat" element={<Chat />} />
                            <Route path="*" element={<NoPage />} />
                        </Route>
                    </Routes>
                </BrowserRouter>
            </React.Fragment>
        </div>
    );
}


// Get initial-state
function init(): IState {
    return {
        sessionData: getEmptySessionData(),
    };
}

// Get blank session-data (avoid undefined errors before login)
function getEmptySessionData(): ISessionData {
    return {
        id: -1,
        email: '',
        name: '',
        waiting: true,
    }
}

// Fetch jwt session-data
async function getSessionData(): Promise<ISessionData> {
    try {
        const data = await authHttp.getSessionData();
        return data;
    } catch (err) {
        console.error(err);
    }
    return getEmptySessionData();
}


// Export default
export default App;
