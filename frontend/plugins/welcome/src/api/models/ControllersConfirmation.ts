/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ControllersConfirmation
 */
export interface ControllersConfirmation {
    /**
     * 
     * @type {string}
     * @memberof ControllersConfirmation
     */
    bookingdate?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersConfirmation
     */
    bookingend?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersConfirmation
     */
    bookingstart?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersConfirmation
     */
    borrow?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersConfirmation
     */
    hourstime?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersConfirmation
     */
    stadium?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersConfirmation
     */
    user?: number;
}

export function ControllersConfirmationFromJSON(json: any): ControllersConfirmation {
    return ControllersConfirmationFromJSONTyped(json, false);
}

export function ControllersConfirmationFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersConfirmation {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bookingdate': !exists(json, 'bookingdate') ? undefined : json['bookingdate'],
        'bookingend': !exists(json, 'bookingend') ? undefined : json['bookingend'],
        'bookingstart': !exists(json, 'bookingstart') ? undefined : json['bookingstart'],
        'borrow': !exists(json, 'borrow') ? undefined : json['borrow'],
        'hourstime': !exists(json, 'hourstime') ? undefined : json['hourstime'],
        'stadium': !exists(json, 'stadium') ? undefined : json['stadium'],
        'user': !exists(json, 'user') ? undefined : json['user'],
    };
}

export function ControllersConfirmationToJSON(value?: ControllersConfirmation | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bookingdate': value.bookingdate,
        'bookingend': value.bookingend,
        'bookingstart': value.bookingstart,
        'borrow': value.borrow,
        'hourstime': value.hourstime,
        'stadium': value.stadium,
        'user': value.user,
    };
}


