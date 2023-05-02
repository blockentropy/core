package keeper

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/v2/x/ml/types"
)

func mlGenerate(model string, prompt string, negprompt string, seed string, signer sdk.AccAddress, blockHeight int64, machine string, endpoint string) {
	fmt.Println("Executing ML Generate Command...")
	// Read the server list from the environment variable
	serverListEnv := os.Getenv("SERVER_LIST")
	if serverListEnv == "" {
		fmt.Println("Error: SERVER_LIST environment variable is not set")
		return
	}

	// Split the comma-separated list of server addresses and ports
	serverList := strings.Split(serverListEnv, ",")

	// Iterate over the list of server addresses and ports
	for _, serverAddr := range serverList {
		// Connect to the server
		conn, err := net.Dial("tcp", serverAddr)
		if err != nil {
			fmt.Printf("Error connecting to the server (%s): %v\n", serverAddr, err)
		} else {
			// Send a message with four arguments
			cmd := "process_generate"
			message := fmt.Sprintf(`%s %s "%s" "%s" %s %s %s "%s" "%s"`, cmd, model, prompt, negprompt, seed, signer.String(), strconv.FormatInt(blockHeight, 10), machine, endpoint)
			_, err = conn.Write([]byte(message))
			if err != nil {
				fmt.Printf("Error sending message to server (%s): %v\n", serverAddr, err)
			}
			defer conn.Close()
		}
	}
}

func (k msgServer) Generate(goCtx context.Context, msg *types.MsgGenerate) (*types.MsgGenerateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	go mlGenerate(msg.Model, msg.Prompt, msg.Negprompt, msg.Seed, msg.GetSigners()[0], ctx.BlockHeight(), msg.Machine, msg.Endpoint)

	return &types.MsgGenerateResponse{}, nil
}
