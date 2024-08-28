package okta

import (
	"io"

	"github.com/okta/okta-cli-client/utils"
	"github.com/spf13/cobra"
)

var HookKeyCmd = &cobra.Command{
	Use:   "hookKey",
	Long:  "Manage HookKeyAPI",
}

func NewHookKeyCmd() *cobra.Command {
    cmd := &cobra.Command{
		Use:   "hookKey",
		Long:  "Manage HookKeyAPI",
	}
	return cmd
}

func init() {
    rootCmd.AddCommand(HookKeyCmd)
}

var (
    
    
            CreateHookKeydata string
        
    
)

func NewCreateHookKeyCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "create",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.HookKeyAPI.CreateHookKey(apiClient.GetConfig().Context)
            
            
            if CreateHookKeydata != "" {
                req = req.Data(CreateHookKeydata)
            }
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
        cmd.Flags().StringVarP(&CreateHookKeydata, "data", "", "", "")
        cmd.MarkFlagRequired("data")
        
    

	return cmd
}

func init() {
	CreateHookKeyCmd := NewCreateHookKeyCmd()
    HookKeyCmd.AddCommand(CreateHookKeyCmd)
}

var (
    
    
    
)

func NewListHookKeysCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "lists",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.HookKeyAPI.ListHookKeys(apiClient.GetConfig().Context)
            
            
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
    

	return cmd
}

func init() {
	ListHookKeysCmd := NewListHookKeysCmd()
    HookKeyCmd.AddCommand(ListHookKeysCmd)
}

var (
    
    
            GetPublicKeypublicKeyId string
        
    
)

func NewGetPublicKeyCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "getPublicKey",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.HookKeyAPI.GetPublicKey(apiClient.GetConfig().Context, GetPublicKeypublicKeyId)
            
            
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
        cmd.Flags().StringVarP(&GetPublicKeypublicKeyId, "publicKeyId", "", "", "")
        cmd.MarkFlagRequired("publicKeyId")
        
    

	return cmd
}

func init() {
	GetPublicKeyCmd := NewGetPublicKeyCmd()
    HookKeyCmd.AddCommand(GetPublicKeyCmd)
}

var (
    
    
            GetHookKeyhookKeyId string
        
    
)

func NewGetHookKeyCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "get",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.HookKeyAPI.GetHookKey(apiClient.GetConfig().Context, GetHookKeyhookKeyId)
            
            
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
        cmd.Flags().StringVarP(&GetHookKeyhookKeyId, "hookKeyId", "", "", "")
        cmd.MarkFlagRequired("hookKeyId")
        
    

	return cmd
}

func init() {
	GetHookKeyCmd := NewGetHookKeyCmd()
    HookKeyCmd.AddCommand(GetHookKeyCmd)
}

var (
    
    
            ReplaceHookKeyhookKeyId string
        
            ReplaceHookKeydata string
        
    
)

func NewReplaceHookKeyCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "replace",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.HookKeyAPI.ReplaceHookKey(apiClient.GetConfig().Context, ReplaceHookKeyhookKeyId)
            
            
            if ReplaceHookKeydata != "" {
                req = req.Data(ReplaceHookKeydata)
            }
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
        cmd.Flags().StringVarP(&ReplaceHookKeyhookKeyId, "hookKeyId", "", "", "")
        cmd.MarkFlagRequired("hookKeyId")
        
        cmd.Flags().StringVarP(&ReplaceHookKeydata, "data", "", "", "")
        cmd.MarkFlagRequired("data")
        
    

	return cmd
}

func init() {
	ReplaceHookKeyCmd := NewReplaceHookKeyCmd()
    HookKeyCmd.AddCommand(ReplaceHookKeyCmd)
}

var (
    
    
            DeleteHookKeyhookKeyId string
        
    
)

func NewDeleteHookKeyCmd() *cobra.Command {
    cmd := &cobra.Command{
	    Use:   "delete",
	  
        RunE: func(cmd *cobra.Command, args []string) error {
            
            
            
            req := apiClient.HookKeyAPI.DeleteHookKey(apiClient.GetConfig().Context, DeleteHookKeyhookKeyId)
            
            
            
            resp, err := req.Execute()
            if err != nil {
                return err
            }
		    d, err := io.ReadAll(resp.Body)
		    if err != nil {
			    return err
		    }
		    utils.PrettyPrintByte(d)
            cmd.Println(string(d))
		    return nil
        },
    }

    
    
        cmd.Flags().StringVarP(&DeleteHookKeyhookKeyId, "hookKeyId", "", "", "")
        cmd.MarkFlagRequired("hookKeyId")
        
    

	return cmd
}

func init() {
	DeleteHookKeyCmd := NewDeleteHookKeyCmd()
    HookKeyCmd.AddCommand(DeleteHookKeyCmd)
}