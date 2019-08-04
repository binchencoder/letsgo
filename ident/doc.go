// Package ident defines the context function to work with user Identity.
//
// User identities are required for systems like Vexillary to calculate the
// launch percentage. Since most backend services would use Vexillary in future,
// they all depend on getting the current userDetails from context.
//
// There are two types of user identities in Jingaol.com. The first type is
// simply called "userDetails" identity which is composed of a company ID, a user
// ID and accout ID. The second type is called "custom" identity which is a string defined
// by applications. Custom identities are used in cases where a normal Jingoal
// user is not required or not existing. For example, the Jingoal enterprise-
// circle application can be accessed by public users who do not have a Jingoal
// account.
//
// To work with a user identity:
//
// ctx := WithUserDetails(context.Background(), userDetails)
// if userDetails, ok := GetUserDetails(ctx); !ok {
// 	fmt.Println("Error, user identity is not set.")
// }
//
// To work with a custom identity:
//
// ctx := WithCustomIdent(context.Background(), "custom string")
// if customId, ok := GetCustomIdent(ctx); !ok {
// 	fmt.Println("Error, custom identity is not set.")
// }
//
package ident
