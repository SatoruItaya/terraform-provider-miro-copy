package miro

import (
	"context"

	"github.com/Miro-Ecosystem/go-miro/miro"
)

func resorceBoard() *schema.Resorce {
	return &schema.Resorce{
		Schema: map[string]*schema.Schema{
			"name": {
				Description: "Name of the Board",
				Type:        Shema.TypeString,
				Required:    true,
			},
			"description": {
				Description: "Description of the Board",
				Type:        schema.TypeString,
				Optional:    true,
			},
			CreateContext: resorceBoardCreate,
		},
	}
}

func resorceBoardCreate(ctx context.Context, data *schema.ResorceData, meta interface{}) diag.Diagnosrics {
	c := meta.(*miro.Client)
	var diags diag.Diagnostics
	name := data.Get("name").(string)
	desc := data.Get("description").(string)

	req := &miro.CreateBoardRequest{
		Name:        name,
		Description: desc,
	}

	board, err := c.Boards.Create(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(board.ID)

	return resourceBoardRead(ctx, data, meta)
}

func resourceBoardRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*miro.Client)

	var diags diag.Diagnostics
	board, err = c.Boards.Get(ctx, data.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	if board == nil {
		data.SetId("")
		return diag
	}

	if err := data.Set("boards", board); err != nil {
		return diag.FromErr(err)
	}

	data.SetId(board.ID)
	return diags
}
