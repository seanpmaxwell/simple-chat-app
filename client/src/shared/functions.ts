/**
 * Convert 'John Paul Smith' to 'JS'.
 */
export function getInitials(name: string): string {
    if (!name) {
        return '';
    }
    const nameArr = name.split(' ');
    const fInitial = nameArr[0].toUpperCase().charAt(0);
    let lInitial = '';
    if (nameArr.length > 1) {
        lInitial = nameArr[nameArr.length - 1].toUpperCase().charAt(0);
    }
    return (fInitial + lInitial);
}
