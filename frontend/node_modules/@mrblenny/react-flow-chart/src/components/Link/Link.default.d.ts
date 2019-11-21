/// <reference types="react" />
import { IConfig, ILink, IOnLinkClick, IOnLinkMouseEnter, IOnLinkMouseLeave, IPosition } from '../../';
export interface ILinkDefaultProps {
    config: IConfig;
    link: ILink;
    startPos: IPosition;
    endPos: IPosition;
    onLinkMouseEnter: IOnLinkMouseEnter;
    onLinkMouseLeave: IOnLinkMouseLeave;
    onLinkClick: IOnLinkClick;
    isHovered: boolean;
    isSelected: boolean;
}
export declare const LinkDefault: ({ config, link, startPos, endPos, onLinkMouseEnter, onLinkMouseLeave, onLinkClick, isHovered, isSelected, }: ILinkDefaultProps) => JSX.Element;
