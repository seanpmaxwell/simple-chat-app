import { useState, useCallback } from 'react';


/**
 * Do setState like react class component.
 */
export function useSetState<T extends object>(
    initialState: T,
): [T, (newPartialState: Partial<T>) => void] {
    const [state, setState] = useState<T>(initialState);
    // Function which accepts a partial state to merge
    const setCustomState = useCallback((newPartialState: Partial<T>) => {
        try {
            setState((prevState): T => {
                return { ...prevState, ...newPartialState };
            });
        } catch (error) {
            console.error(error);
        }
    }, []);
    return [state, setCustomState];
};
