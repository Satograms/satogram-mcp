package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerTools(srv *server.MCPServer, client *SatogramClient) {
	// 1. get_recipient_stats
	srv.AddTool(mcp.NewTool(
		"get_recipient_stats",
		mcp.WithDescription("Get statistics about Satogram recipients: how many node pubkeys, lightning addresses, total recipients, satograms sent, and signups."),
	), handleGetRecipientStats(client))

	// 2. create_satogram
	srv.AddTool(mcp.NewTool(
		"create_satogram",
		mcp.WithDescription("Create a new Satogram campaign. Sends a message with sats to Lightning Network nodes. Returns a Lightning invoice (payment_request) that must be paid to start the campaign."),
		mcp.WithNumber("total_cost",
			mcp.Description("Total sats to spend on this campaign (500-250000)."),
			mcp.Required(),
		),
		mcp.WithString("message",
			mcp.Description("The message to send with each keysend payment."),
			mcp.Required(),
		),
		mcp.WithNumber("amt_per_satogram",
			mcp.Description("Sats to send per recipient (default: 1)."),
		),
		mcp.WithNumber("max_fees",
			mcp.Description("Max routing fee per payment in sats (default: 20)."),
		),
		mcp.WithString("sender_address",
			mcp.Description("Optional lightning address to include as sender."),
		),
		mcp.WithString("recipient_selection",
			mcp.Description("Which recipients to target: 'all' or 'signups' (default: 'all')."),
			mcp.Enum("all", "signups"),
		),
	), handleCreateSatogram(client))

	// 3. check_invoice_status
	srv.AddTool(mcp.NewTool(
		"check_invoice_status",
		mcp.WithDescription("Check the payment status of a Satogram invoice. Returns OPEN (unpaid), SETTLED (paid), CANCELED, or ACCEPTED."),
		mcp.WithString("payment_request",
			mcp.Description("The Lightning invoice (payment_request) returned by create_satogram."),
			mcp.Required(),
		),
	), handleCheckInvoiceStatus(client))

	// 4. get_satogram_status
	srv.AddTool(mcp.NewTool(
		"get_satogram_status",
		mcp.WithDescription("Get the delivery status of a Satogram campaign. Shows success/failure counts, sats spent, and overall progress. Status values: 0=Completed, 1=InProgress, 2=NotStarted."),
		mcp.WithString("payment_request",
			mcp.Description("The Lightning invoice (payment_request) returned by create_satogram."),
			mcp.Required(),
		),
	), handleGetSatogramStatus(client))
}

func handleGetRecipientStats(client *SatogramClient) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		stats, err := client.GetRecipientStats()
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get recipient stats: %v", err)), nil
		}

		data, err := json.MarshalIndent(stats, "", "  ")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to format response: %v", err)), nil
		}
		return mcp.NewToolResultText(string(data)), nil
	}
}

func handleCreateSatogram(client *SatogramClient) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		totalCost, ok := args["total_cost"].(float64)
		if !ok {
			return mcp.NewToolResultError("total_cost is required and must be a number"), nil
		}
		message, ok := args["message"].(string)
		if !ok || message == "" {
			return mcp.NewToolResultError("message is required and must be a non-empty string"), nil
		}

		payload := &SatogramPayload{
			TotalCost: int64(totalCost),
			Message:   message,
		}

		if v, ok := args["amt_per_satogram"].(float64); ok {
			amt := int64(v)
			payload.AmtPerSatogram = &amt
		}
		if v, ok := args["max_fees"].(float64); ok {
			mf := int64(v)
			payload.MaxFees = &mf
		}
		if v, ok := args["sender_address"].(string); ok && v != "" {
			payload.SenderAddress = &v
		}
		if v, ok := args["recipient_selection"].(string); ok && v != "" {
			payload.RecipientSelection = &v
		}

		result, err := client.CreateSatogram(payload)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create satogram: %v", err)), nil
		}

		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to format response: %v", err)), nil
		}
		return mcp.NewToolResultText(string(data)), nil
	}
}

func handleCheckInvoiceStatus(client *SatogramClient) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		paymentRequest, ok := args["payment_request"].(string)
		if !ok || paymentRequest == "" {
			return mcp.NewToolResultError("payment_request is required"), nil
		}

		status, err := client.CheckInvoiceStatus(paymentRequest)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to check invoice status: %v", err)), nil
		}

		data, err := json.MarshalIndent(status, "", "  ")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to format response: %v", err)), nil
		}
		return mcp.NewToolResultText(string(data)), nil
	}
}

func handleGetSatogramStatus(client *SatogramClient) server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		paymentRequest, ok := args["payment_request"].(string)
		if !ok || paymentRequest == "" {
			return mcp.NewToolResultError("payment_request is required"), nil
		}

		stored, err := client.GetSatogramStatus(paymentRequest)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get satogram status: %v", err)), nil
		}

		data, err := json.MarshalIndent(stored, "", "  ")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to format response: %v", err)), nil
		}
		return mcp.NewToolResultText(string(data)), nil
	}
}
